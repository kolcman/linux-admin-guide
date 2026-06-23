package data

var FirewallSection = Section{
	Title: "🔥 UFW И IPTABLES — УПРАВЛЕНИЕ ПРАВИЛАМИ",
	Items: []Item{
		{Type: TypeHeader, Value: "✅ Как РАЗРЕШАТЬ (Allow):"},
		{Type: TypeCmd, Value: "sudo ufw allow 80/tcp", Desc: "Разрешить TCP трафик на порт 80"},
		{Type: TypeCmd, Value: "sudo ufw allow from 192.168.1.100", Desc: "Разрешить ВСЕ порты только с этого IP"},
		{Type: TypeCmd, Value: "sudo ufw allow from 10.0.0.0/24 to any port 3306", Desc: "Разрешить MySQL только из подсети"},
		{Type: TypeCmd, Value: "sudo ufw limit 22/tcp", Desc: "Разрешить SSH с защитой от брутфорса"},

		{Type: TypeHeader, Value: "⛔ Как ЗАПРЕЩАТЬ (Deny/Reject):"},
		{Type: TypeCmd, Value: "sudo ufw deny 8080", Desc: "Запретить доступ к порту 8080 (молча)"},
		{Type: TypeCmd, Value: "sudo ufw reject 25", Desc: "Отклонить порт 25 (с уведомлением отправителя)"},
		{Type: TypeCmd, Value: "sudo ufw deny from 1.2.3.4", Desc: "Полный бан конкретного IP адреса"},

		{Type: TypeHeader, Value: "🗑️ Как УДАЛЯТЬ правила:"},
		{Type: TypeCmd, Value: "sudo ufw status numbered", Desc: "1. Посмотреть список с номерами"},
		{Type: TypeCmd, Value: "sudo ufw delete 3", Desc: "2. Удалить правило под номером 3 из списка"},
		{Type: TypeCmd, Value: "sudo ufw delete allow 80/tcp", Desc: "3. Удалить по описанию (менее надежно)"},
		{Type: TypeCmd, Value: "sudo ufw reset", Desc: "4. Сбросить ВСЕ правила к заводским настройкам"},

		{Type: TypeHeader, Value: "🔍 Диагностика через iptables:"},
		{Type: TypeCmd, Value: "sudo iptables -L -n -v", Desc: "Детальная статистика пакетов (без DNS)"},
		{Type: TypeCmd, Value: "sudo iptables -Z", Desc: "Сбросить счетчики пакетов в ноль"},

		{Type: TypeHeader, Value: "⚙️ Важные настройки:"},
		{Type: TypeCmd, Value: "sudo ufw default deny incoming", Desc: "Политика: запрещать всё входящее по умолчанию"},
		{Type: TypeCmd, Value: "sudo ufw enable", Desc: "Включить файрвол (подтвердите действие)"},
	},
}
