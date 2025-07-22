// Package errs реализует очень простой механизм единообразного формата ошибок, возвращаемых web-api.
package errs

var delimiter = "."

// SetDelimiter устанавливает используемый при слиянии Errors разделитель между ключом добавляемой ошибки
// и приписываемым слева к этому ключу префиксом.
func SetDelimiter(s string) {
	delimiter = s
}

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

// AddErrors добавляет информацию из хранилища add к хранилищу err.
// Если добавляемый ключ - пустая строка, в качестве нового ключа используется префикс prefix.
// Иначе к ключу слева дописывается префикс+разделитель и полученная строка используется как новый ключ.
//
// Это позволяет сформировать удобный для обработки путь к ошибочному элементу, состоящий из имён полей и индексов массивов.
func (err *Errors) AddErrors(prefix string, add Errors) *Errors {
	for key, val := range add {
		if key == "" {
			err.Add(prefix, val...)
		} else {
			err.Add(prefix+delimiter+key, val...)
		}
	}
	return err
}
