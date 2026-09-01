package data

// UsersCreateSection — создание учётки, пароль, срок жизни.
var UsersCreateSection = Section{
	Title: "👤 СОЗДАНИЕ УЧЁТКИ",
	Items: []Item{
		{Type: TypeTip, Value: "UID — числовой ID учётки. Root = 0 (полные права).\nСистемные службы: обычно низкие UID. Обычные пользователи: часто от 1000.\nДомашний каталог — личная папка (/home/имя): данные и настройки."},

		{Type: TypeHeader, Value: "👤 adduser (Debian / Ubuntu)"},
		{Type: TypeTip, Value: "adduser — удобная обёртка: часто спросит пароль и создаст home сама.\nНа RHEL/Fedora adduser часто = useradd — там лучше useradd + passwd."},
		{Type: TypeKey, Key: "--home КАТ", Desc: "путь домашнего каталога"},
		{Type: TypeKey, Key: "--shell ОБОЛОЧКА", Desc: "login shell, например /bin/bash"},
		{Type: TypeKey, Key: "--no-create-home", Desc: "не создавать домашний каталог"},
		{Type: TypeKey, Key: "--disabled-login", Desc: "учётка без обычного входа"},
		{Type: TypeKey, Key: "--disabled-password", Desc: "без пароля (ключи и т.п. — по политике)"},
		{Type: TypeCmd, Value: "sudo adduser alice", Desc: "создать пользователя (интерактивно)"},
		{Type: TypeCmd, Value: "sudo adduser --home /home/alice --shell /bin/bash alice", Desc: "с явным home и shell"},
		{Type: TypeCmd, Value: "sudo adduser --disabled-login svcbot", Desc: "служебная учётка без логина"},

		{Type: TypeHeader, Value: "🔑 useradd (везде)"},
		{Type: TypeKey, Key: "-m", Desc: "создать домашнюю папку (/home/user)"},
		{Type: TypeKey, Key: "-s /bin/bash", Desc: "указать оболочку (по умолчанию часто sh)"},
		{Type: TypeKey, Key: "-g <group>", Desc: "задать основную группу"},
		{Type: TypeKey, Key: "-G <group1,group2>", Desc: "добавить в дополнительные группы"},
		{Type: TypeCmd, Value: "sudo useradd -m -s /bin/bash username", Desc: "создать пользователя с home и bash"},
		{Type: TypeCmd, Value: "sudo passwd username", Desc: "задать пароль (useradd сам не спросит)"},
		{Type: TypeCmd, Value: "sudo userdel -r username", Desc: "удалить пользователя и его домашнюю папку"},

		{Type: TypeHeader, Value: "🔐 passwd"},
		{Type: TypeKey, Key: "-S", Desc: "статус пароля учётки"},
		{Type: TypeKey, Key: "-e", Desc: "просрочить пароль — сменить при следующем входе"},
		{Type: TypeKey, Key: "-d", Desc: "удалить пароль (осторожно)"},
		{Type: TypeKey, Key: "-n / -x / -w", Desc: "min / max дней жизни пароля / дней предупреждения"},
		{Type: TypeKey, Key: "-i", Desc: "дней после просрочки до неактивности учётки"},
		{Type: TypeCmd, Value: "sudo passwd -S alice", Desc: "статус пароля"},
		{Type: TypeCmd, Value: "sudo passwd -e alice", Desc: "потребовать смену пароля при входе"},
		{Type: TypeCmd, Value: "sudo passwd -n 7 -x 90 -w 7 alice", Desc: "min 7 / max 90 дней, warn за 7"},

		{Type: TypeHeader, Value: "📅 chage — срок пароля"},
		{Type: TypeKey, Key: "-l", Desc: "показать текущую политику"},
		{Type: TypeKey, Key: "-M / -W", Desc: "max дней жизни / дней предупреждения"},
		{Type: TypeKey, Key: "-E ДАТА", Desc: "дата истечения учётки (YYYY-MM-DD)"},
		{Type: TypeKey, Key: "-I", Desc: "дней неактивности после просрочки пароля"},
		{Type: TypeCmd, Value: "sudo chage -l alice", Desc: "политика пароля alice"},
		{Type: TypeCmd, Value: "sudo chage -M 90 -W 7 alice", Desc: "пароль max 90 дней, warn за 7"},
		{Type: TypeCmd, Value: "sudo chage -E 2026-12-31 alice", Desc: "учётка до указанной даты"},
		{Type: TypeCmd, Value: "sudo chage -I 14 alice", Desc: "14 дней неактивности после просрочки пароля"},
	},
}

// UsersGroupsSection — группы, usermod, информация.
var UsersGroupsSection = Section{
	Title: "👥 ГРУППЫ И ИНФОРМАЦИЯ",
	Items: []Item{
		{Type: TypeHeader, Value: "⚙️ usermod"},
		{Type: TypeKey, Key: "-aG", Desc: "добавить в группу, не затирая остальные"},
		{Type: TypeKey, Key: "-L / -U", Desc: "заблокировать / разблокировать учётку"},
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
		{Type: TypeCmd, Value: "groups username", Desc: "список групп"},
		{Type: TypeCmd, Value: "getent passwd username", Desc: "строка учётки"},
		{Type: TypeCmd, Value: "ls -la /home/username", Desc: "проверить домашний каталог"},
		{Type: TypeTip, Value: "После смены групп (usermod -aG) перелогинься — иначе новая группа не подхватится в текущей сессии."},
	},
}

// UsersSudoSection — su, sudo, краткий chmod/chown.
var UsersSudoSection = Section{
	Title: "🛡️ ДЕЛЕГИРОВАНИЕ ПРАВ",
	Items: []Item{
		{Type: TypeTip, Value: "Не сиди постоянно под root. su — стать другим пользователем.\nsudo — выполнить команду с привилегиями по правилам sudoers.\nchmod/chown — права и владелец файлов (подробнее: Файлы → Права)."},

		{Type: TypeHeader, Value: "🔄 su — сменить пользователя"},
		{Type: TypeCmd, Value: "su", Desc: "переключиться на root (нужен пароль root)"},
		{Type: TypeCmd, Value: "su -", Desc: "root + login-окружение (рекомендуемый вид)"},
		{Type: TypeCmd, Value: "su - alice", Desc: "войти как alice с её окружением"},
		{Type: TypeCmd, Value: "su -c \"id\" alice", Desc: "одна команда от имени alice"},
		{Type: TypeTip, Value: "Если пароль root отключён — su на root не сработает, пользуйся sudo."},

		{Type: TypeHeader, Value: "🛡️ sudo"},
		{Type: TypeTip, Value: "sudo даёт определённым пользователям запускать указанные команды\nс админ-правами — обычно со своим паролем, без пароля root."},
		{Type: TypeCmd, Value: "sudo apt update", Desc: "выполнить команду от root"},
		{Type: TypeCmd, Value: "sudo -u www-data id", Desc: "выполнить от имени другого пользователя"},
		{Type: TypeCmd, Value: "sudo -i", Desc: "интерактивный root-shell"},
		{Type: TypeCmd, Value: "sudo -l", Desc: "что разрешено текущему пользователю"},
		{Type: TypeCmd, Value: "sudo usermod -aG sudo alice", Desc: "Debian/Ubuntu: группа sudo"},
		{Type: TypeCmd, Value: "sudo usermod -aG wheel alice", Desc: "RHEL/Fedora: часто wheel"},
		{Type: TypeCmd, Value: "sudo visudo", Desc: "безопасно править /etc/sudoers"},
		{Type: TypeTip, Value: "sudoers:\nalice ALL=(ALL:ALL) ALL\n— alice может всё через sudo.\n\nalice ALL=(ALL) NOPASSWD: /bin/systemctl restart nginx\n— только эта команда без пароля."},
		{Type: TypeWarn, Value: "Не правь /etc/sudoers обычным редактором вслепую — опечатка закроет sudo.\nТолько visudo."},

		{Type: TypeHeader, Value: "🔐 chmod / chown"},
		{Type: TypeTip, Value: "Права файлов подробно: Файлы → Права и маски."},
		{Type: TypeCmd, Value: "chmod 755 script.sh", Desc: "владелец всё, остальные читают и запускают"},
		{Type: TypeCmd, Value: "sudo chown alice:alice file.txt", Desc: "владелец и группа"},
		{Type: TypeCmd, Value: "sudo chown -R www-data:www-data /var/www", Desc: "рекурсивно отдать дерево"},
		{Type: TypeCmd, Value: "sudo chown -h alice:alice link", Desc: "сменить владельца symlink, не цели"},
		{Type: TypeKey, Key: "-R", Desc: "chmod/chown рекурсивно"},
		{Type: TypeKey, Key: "-h", Desc: "chown: менять symlink, а не цель"},
	},
}

// UsersAclSection — ACL, default ACL (наследование), umask.
var UsersAclSection = Section{
	Title: "📎 ACL И UMASK",
	Items: []Item{
		{Type: TypeTip, Value: "Обычных u/g/o часто мало. ACL — права конкретным людям и группам.\nDefault ACL на каталоге = наследование для новых файлов/папок внутри.\numask — какие права получат новые объекты при создании."},

		{Type: TypeHeader, Value: "👀 getfacl — посмотреть"},
		{Type: TypeCmd, Value: "getfacl test.txt", Desc: "ACL файла"},
		{Type: TypeCmd, Value: "getfacl *", Desc: "все объекты в текущем каталоге"},
		{Type: TypeCmd, Value: "getfacl -R /shared", Desc: "рекурсивно"},

		{Type: TypeHeader, Value: "✏️ setfacl — задать"},
		{Type: TypeKey, Key: "-m", Desc: "модифицировать ACL (остальное сохранить)"},
		{Type: TypeKey, Key: "-x", Desc: "удалить указанные ACL-записи"},
		{Type: TypeKey, Key: "--set", Desc: "заменить весь набор ACL"},
		{Type: TypeKey, Key: "-b", Desc: "снять все ACL, оставить базовые права"},
		{Type: TypeKey, Key: "-k", Desc: "удалить default ACL с каталога"},
		{Type: TypeKey, Key: "-d", Desc: "работать с ACL по умолчанию (наследование)"},
		{Type: TypeKey, Key: "-R", Desc: "рекурсивно"},
		{Type: TypeKey, Key: "u:user:rwx", Desc: "правило для пользователя"},
		{Type: TypeKey, Key: "g:group:rx", Desc: "правило для группы"},
		{Type: TypeKey, Key: "m:rx", Desc: "маска эффективных прав"},
		{Type: TypeKey, Key: "o:---", Desc: "остальные"},
		{Type: TypeCmd, Value: "setfacl -m u:alice:rw file.txt", Desc: "alice — чтение и запись"},
		{Type: TypeCmd, Value: "setfacl -m g:devs:rx file.txt", Desc: "группе devs — r и x"},
		{Type: TypeCmd, Value: "setfacl -m m:rx file.txt", Desc: "маска эффективных прав"},
		{Type: TypeCmd, Value: "setfacl -m d:u:alice:rwx /shared", Desc: "default: новые объекты дадут alice rwx"},
		{Type: TypeCmd, Value: "setfacl -b file.txt", Desc: "убрать расширенные ACL"},
		{Type: TypeTip, Value: "Наследование = default ACL (-d или d:…) на каталоге.\nУже лежащие файлы сами не обновятся — им ACL ставят отдельно (часто -R)."},
		{Type: TypeWarn, Value: "Нужна ФС с ACL и пакет acl (getfacl/setfacl)."},

		{Type: TypeHeader, Value: "🎭 umask"},
		{Type: TypeTip, Value: "База при создании: файлы 666, каталоги 777. umask вычитает биты.\numask 022 → файлы 644, каталоги 755."},
		{Type: TypeCmd, Value: "umask", Desc: "текущая маска (часто 0022)"},
		{Type: TypeCmd, Value: "umask -S", Desc: "символьно: u=rwx,g=rx,o=rx"},
		{Type: TypeCmd, Value: "umask 027", Desc: "файлы ~640, каталоги ~750"},
		{Type: TypeTip, Value: "umask — на текущий shell. Постоянно: ~/.bashrc или системный профиль."},
	},
}

func usersSubmenu() []MenuItem {
	return []MenuItem{
		{Title: UsersCreateSection.Title, Section: &UsersCreateSection},
		{Title: UsersGroupsSection.Title, Section: &UsersGroupsSection},
		{Title: UsersSudoSection.Title, Section: &UsersSudoSection},
		{Title: UsersAclSection.Title, Section: &UsersAclSection},
	}
}
