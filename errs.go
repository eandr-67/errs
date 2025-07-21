// Package errs реализует очень простой механизм единообразного формата ошибок, возвращаемых web-api.
package errs

// Errors определяет тип хранилища ошибок:
// ключ идентифицирует сущность, к которой относятся ошибки,
// массив строк - информация об ошибках, относящихся к этой сущности.
type Errors map[string][]string

// Add добавляет ошибки с текстами msg к сущности с ключом key
func (err *Errors) Add(key string, msg ...string) *Errors {
	if len(msg) == 0 {
		return err
	}
	if (*err) == nil {
		*err = Errors{key: msg}
	} else {
		(*err)[key] = append((*err)[key], msg...)
	}
	return err
}

// AddErrors добавляет информацию из хранилища err к хранилищу e.
// Если добавляемый ключ - пустая строка или ключ начинается символом '[', к нему слева дописывается prefix.
// Иначе к ключу слева дописывается prefix+".".
//
// Это позволяет сформировать корректный JS-путь к ошибочному элементу, состоящий из имён полей и индексов массивов.
func (err *Errors) AddErrors(prefix string, add Errors) *Errors {
	for key, val := range add {
		if key == "" || key[0] == '[' {
			err.Add(prefix+key, val...)
		} else {
			err.Add(prefix+"."+key, val...)
		}
	}
	return err
}
