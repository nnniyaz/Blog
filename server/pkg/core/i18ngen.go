// code generated by gen.go
// DO NOT EDIT.
package core

const (
	TXT_UNKNOWN_ERROR                      TxtKey = 1
	TXT_UNAUTHORIZED                       TxtKey = 2
	TXT_WRONG_MLSTRING_FORMAT              TxtKey = 3
	TXT_UUID_INVALID_LENGTH                TxtKey = 4
	TXT_EMPTY_PASSWORD                     TxtKey = 5
	TXT_PASSWORD_TOO_SHORT                 TxtKey = 6
	TXT_PASSWORD_TOO_LONG                  TxtKey = 7
	TXT_INVALID_EMAIL                      TxtKey = 8
	TXT_ARTICLES_TITLE_IS_EMPTY            TxtKey = 9
	TXT_ARTICLES_TITLE_IS_TOO_LONG         TxtKey = 10
	TXT_ARTICLES_CONTENT_IS_EMPTY          TxtKey = 11
	TXT_NAME_OF_CONTACT_OR_SOCIAL_IS_EMPTY TxtKey = 12
	TXT_FIRST_NAME_IS_EMPTY                TxtKey = 13
	TXT_LAST_NAME_IS_EMPTY                 TxtKey = 14
	TXT_BIO_IS_EMPTY                       TxtKey = 15
	TXT_TITLE_OF_BOOK_IS_EMPTY             TxtKey = 16
	TXT_AUTHOR_OF_BOOK_IS_EMPTY            TxtKey = 17
	TXT_NAME_OF_PROJECT_IS_EMPTY           TxtKey = 18
	TXT_USER_AGENT_IS_EMPTY                TxtKey = 19
)

var Txts = TxtResource{
	TXT_UNKNOWN_ERROR: MlString{
		KZ: `Белгісіз қате`,
		RU: `Неизвестная ошибка`,
		EN: `Unknown error`,
	},
	TXT_UNAUTHORIZED: MlString{
		KZ: `Құқығы жоқ`,
		RU: `Нет доступа`,
		EN: `Unauthorized`,
	},
	TXT_WRONG_MLSTRING_FORMAT: MlString{
		KZ: `Мәтіндің түрі дұрыс емес`,
		RU: `Неверный формат текста`,
		EN: `Wrong mlstring format`,
	},
	TXT_UUID_INVALID_LENGTH: MlString{
		KZ: `UUID ұзындығы дұрыс емес`,
		RU: `Неверная длина UUID`,
		EN: `UUID invalid length`,
	},
	TXT_EMPTY_PASSWORD: MlString{
		KZ: `Құпия сөз бос`,
		RU: `Пароль пустой`,
		EN: `Empty password`,
	},
	TXT_PASSWORD_TOO_SHORT: MlString{
		KZ: `Құпия сөз ұзындығы тым қысқа`,
		RU: `Пароль слишком короткий`,
		EN: `Password too short`,
	},
	TXT_PASSWORD_TOO_LONG: MlString{
		KZ: `Құпия сөз ұзындығы тым ұзын`,
		RU: `Пароль слишком длинный`,
		EN: `Password too long`,
	},
	TXT_INVALID_EMAIL: MlString{
		KZ: `Email дұрыс емес`,
		RU: `Неверный email`,
		EN: `Invalid email`,
	},
	TXT_ARTICLES_TITLE_IS_EMPTY: MlString{
		KZ: `Мақала тақырыбы бос`,
		RU: `Заголовок статьи пуст`,
		EN: `Article title is empty`,
	},
	TXT_ARTICLES_TITLE_IS_TOO_LONG: MlString{
		KZ: `Мақаланың тақырыбы тым ұзын`,
		RU: `Заголовок статьи слишком длинный`,
		EN: `Article title is too long`,
	},
	TXT_ARTICLES_CONTENT_IS_EMPTY: MlString{
		KZ: `Мақаланың мазмұны бос`,
		RU: `Содержимое статьи пустое`,
		EN: `Article content is empty`,
	},
	TXT_NAME_OF_CONTACT_OR_SOCIAL_IS_EMPTY: MlString{
		KZ: `Контакт немесе социалдық ақпараттың аты бос`,
		RU: `Название контакта или социальной сети пустое`,
		EN: `Name of contact or social is empty`,
	},
	TXT_FIRST_NAME_IS_EMPTY: MlString{
		KZ: `Аты бос`,
		RU: `Имя пустое`,
		EN: `First name is empty`,
	},
	TXT_LAST_NAME_IS_EMPTY: MlString{
		KZ: `Тегі бос`,
		RU: `Фамилия пустая`,
		EN: `Last name is empty`,
	},
	TXT_BIO_IS_EMPTY: MlString{
		KZ: `Биография бос`,
		RU: `Биография пустая`,
		EN: `Bio is empty`,
	},
	TXT_TITLE_OF_BOOK_IS_EMPTY: MlString{
		KZ: `Кітаптың атауы бос`,
		RU: `Название книги пустое`,
		EN: `Book title is empty`,
	},
	TXT_AUTHOR_OF_BOOK_IS_EMPTY: MlString{
		KZ: `Кітаптың авторы бос`,
		RU: `Автор книги пустой`,
		EN: `Author of book is empty`,
	},
	TXT_NAME_OF_PROJECT_IS_EMPTY: MlString{
		KZ: `Проекттің аты бос`,
		RU: `Название проекта пустое`,
		EN: `Name of project is empty`,
	},
	TXT_USER_AGENT_IS_EMPTY: MlString{
		KZ: `User agent бос`,
		RU: `User agent пустой`,
		EN: `User agent is empty`,
	},
}

func GetTxtKeyAsString(k TxtKey) string {
	switch k {
	case TXT_UNKNOWN_ERROR:
		return "unknown_error"
	case TXT_UNAUTHORIZED:
		return "unauthorized"
	case TXT_WRONG_MLSTRING_FORMAT:
		return "wrong_mlstring_format"
	case TXT_UUID_INVALID_LENGTH:
		return "uuid_invalid_length"
	case TXT_EMPTY_PASSWORD:
		return "empty_password"
	case TXT_PASSWORD_TOO_SHORT:
		return "password_too_short"
	case TXT_PASSWORD_TOO_LONG:
		return "password_too_long"
	case TXT_INVALID_EMAIL:
		return "invalid_email"
	case TXT_ARTICLES_TITLE_IS_EMPTY:
		return "articles_title_is_empty"
	case TXT_ARTICLES_TITLE_IS_TOO_LONG:
		return "articles_title_is_too_long"
	case TXT_ARTICLES_CONTENT_IS_EMPTY:
		return "articles_content_is_empty"
	case TXT_NAME_OF_CONTACT_OR_SOCIAL_IS_EMPTY:
		return "name_of_contact_or_social_is_empty"
	case TXT_FIRST_NAME_IS_EMPTY:
		return "first_name_is_empty"
	case TXT_LAST_NAME_IS_EMPTY:
		return "last_name_is_empty"
	case TXT_BIO_IS_EMPTY:
		return "bio_is_empty"
	case TXT_TITLE_OF_BOOK_IS_EMPTY:
		return "title_of_book_is_empty"
	case TXT_AUTHOR_OF_BOOK_IS_EMPTY:
		return "author_of_book_is_empty"
	case TXT_NAME_OF_PROJECT_IS_EMPTY:
		return "name_of_project_is_empty"
	case TXT_USER_AGENT_IS_EMPTY:
		return "user_agent_is_empty"

	default:
		return ""
	}
}
