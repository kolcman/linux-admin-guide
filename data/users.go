package data

var UsersSection = Section{
	Title: "👤 ПОЛЬЗОВАТЕЛИ И ГРУППЫ",
	Items: []Item{
		{Type: TypeHeader, Value: "🔑 Ключи useradd (создание):"},
		{Type: TypeKey, Key: "-m", Desc: "Создать домашнюю папку (/home/user)"},
		{Type: TypeKey, Key: "-s /bin/bash", Desc: "Указать оболочку (по умолч. sh)"},
		{Type: TypeKey, Key: "-g <group>", Desc: "Задать основную группу"},
		{Type: TypeKey, Key: "-G <group1,group2>", Desc: "Добавить в дополнительные группы"},

		{Type: TypeHeader, Value: "🛠️ Основные команды:"},
		{Type: TypeCmd, Value: "sudo useradd -m -s /bin/bash username", Desc: "Создать нового пользователя"},
		{Type: TypeCmd, Value: "sudo passwd username", Desc: "Задать или сменить пароль"},
		{Type: TypeCmd, Value: "sudo userdel -r username", Desc: "Удалить пользователя и его домашнюю папку"},

		{Type: TypeHeader, Value: "⚙️ Управление правами (usermod):"},
		{Type: TypeCmd, Value: "sudo usermod -aG sudo username", Desc: "Дать права администратора (sudo)"},
		{Type: TypeCmd, Value: "sudo usermod -aG docker username", Desc: "Дать доступ к Docker без sudo"},
		{Type: TypeCmd, Value: "sudo usermod -L username", Desc: "Заблокировать учетную запись (Lock)"},

		{Type: TypeHeader, Value: "📋 Информация и группы:"},
		{Type: TypeCmd, Value: "id username", Desc: "Показать UID, GID и группы пользователя"},
		{Type: TypeCmd, Value: "groups username", Desc: "Список групп, в которых состоит пользователь"},
		{Type: TypeCmd, Value: "cat /etc/passwd | grep username", Desc: "Проверить данные в системном файле"},
	},
}
