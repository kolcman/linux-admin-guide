package data

var UsersSection = Section{
	Title: "👤 ПОЛЬЗОВАТЕЛИ И ГРУППЫ",
	Items: []Item{
		{Type: TypeTip, Value: "UID — числовой ID учётки. Root = 0 (полные права).\nСистемные службы: обычно 1–999. Обычные пользователи: от 1000."},

		{Type: TypeHeader, Value: "🔑 Ключи useradd (создание)"},
		{Type: TypeKey, Key: "-m", Desc: "создать домашнюю папку (/home/user)"},
		{Type: TypeKey, Key: "-s /bin/bash", Desc: "указать оболочку (по умолчанию часто sh)"},
		{Type: TypeKey, Key: "-g <group>", Desc: "задать основную группу"},
		{Type: TypeKey, Key: "-G <group1,group2>", Desc: "добавить в дополнительные группы"},

		{Type: TypeHeader, Value: "🛠️ Основные команды"},
		{Type: TypeCmd, Value: "sudo useradd -m -s /bin/bash username", Desc: "создать нового пользователя"},
		{Type: TypeCmd, Value: "sudo passwd username", Desc: "задать или сменить пароль"},
		{Type: TypeCmd, Value: "sudo userdel -r username", Desc: "удалить пользователя и его домашнюю папку"},

		{Type: TypeHeader, Value: "⚙️ Управление правами (usermod)"},
		{Type: TypeCmd, Value: "sudo usermod -aG sudo username", Desc: "дать права администратора (sudo)"},
		{Type: TypeCmd, Value: "sudo usermod -aG docker username", Desc: "дать доступ к Docker без sudo"},
		{Type: TypeCmd, Value: "sudo usermod -L username", Desc: "заблокировать учётку (Lock)"},
		{Type: TypeCmd, Value: "sudo usermod -U username", Desc: "разблокировать учётку"},
		{Type: TypeWarn, Value: "usermod -G без -a перезаписывает ВЕСЬ список доп. групп — легко снести sudo/docker.\nВсегда: usermod -aG."},

		{Type: TypeHeader, Value: "👥 Группы"},
		{Type: TypeCmd, Value: "sudo groupadd groupname", Desc: "создать группу"},
		{Type: TypeCmd, Value: "sudo groupdel groupname", Desc: "удалить группу (пустую / без primary)"},
		{Type: TypeCmd, Value: "sudo gpasswd -a username groupname", Desc: "добавить пользователя в группу"},
		{Type: TypeCmd, Value: "sudo gpasswd -d username groupname", Desc: "убрать пользователя из группы"},
		{Type: TypeCmd, Value: "getent group groupname", Desc: "состав группы: имя и участники"},

		{Type: TypeHeader, Value: "📋 Информация"},
		{Type: TypeCmd, Value: "id username", Desc: "UID, GID и группы пользователя"},
		{Type: TypeCmd, Value: "groups username", Desc: "список групп, в которых состоит пользователь"},
		{Type: TypeCmd, Value: "cat /etc/passwd | grep username", Desc: "проверить данные в системном файле"},
		{Type: TypeTip, Value: "После смены групп (usermod -aG) перелогинься — иначе новая группа не подхватится в текущей сессии."},
	},
}
