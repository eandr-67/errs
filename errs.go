// Package errs реализует очень простой механизм единообразного формата ошибок, возвращаемых web-api.
package errs

// Errors определяет тип хранилища ошибок:
// ключ идентифицирует сущность, к которой относятся ошибки,
// массив строк - информация об ошибках, относящихся к этой сущности.
type Errors map[string][]string

// Add добавляет ошибки с текстами messages к сущности с ключом key
func (e *Errors) Add(key string, messages ...string) {
	if len(messages) == 0 {
		return
	}
	if v, ok := (*e)[key]; !ok {
		(*e)[key] = messages
	} else {
		(*e)[key] = append(v, messages...)
	}
}

// AddErrors добавляет информацию из хранилища err к хранилищу e.
// К каждому ключу, полученному из err, перед добавлением слева приписывается prefix.
func (e *Errors) AddErrors(prefix string, errors Errors) {
	for k, v := range errors {
		e.Add(prefix+k, v...)
	}
}
