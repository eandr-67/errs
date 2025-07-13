// Package errs реализует очень простой механизм единообразного формата ошибок, возвращаемых web-api.
package errs

// Errors определяет тип хранилища ошибок:
// ключ идентифицирует сущность, к которой относятся ошибки,
// массив строк - информация об ошибках, относящихся к этой сущности.
type Errors map[string][]string

// Add добавляет ошибку с текстом message к сущности с ключом key
func (e *Errors) Add(key string, message string) {
	e.initIfNil()
	if _, ok := (*e)[key]; !ok {
		(*e)[key] = []string{message}
	} else {
		(*e)[key] = append((*e)[key], message)
	}
}

// AddMessages добавляет массив текстов ошибок messages к сущности с ключом key
func (e *Errors) AddMessages(key string, messages []string) {
	e.initIfNil()
	e.addMessages(key, messages)
}

// AddErrors добавляет информацию из хранилища err к хранилищу e.
// К каждому ключу, полученному из err, перед добавлением слева приписывается prefix.
func (e *Errors) AddErrors(prefix string, err Errors) {
	e.initIfNil()
	for k, v := range err {
		e.addMessages(prefix+k, v)
	}
}

// addMessages вспомогательный метод, используемый в AddMassages и AddErrors для добавления массива текстов ошибок.
// Нужен для исключения лишних вызовов initIfNil при выполнении AddErrors.
func (e *Errors) addMessages(key string, messages []string) {
	if len(messages) == 0 {
		return
	}
	if _, ok := (*e)[key]; !ok {
		(*e)[key] = messages
	} else {
		(*e)[key] = append((*e)[key], messages...)
	}
}

// initIfNil вспомогательный метод, инициализирующий хранилище ошибок, если был передан указатель на nil.
func (e *Errors) initIfNil() {
	if *e == nil {
		*e = map[string][]string{}
	}
}
