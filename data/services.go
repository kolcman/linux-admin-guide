package data

var ServicesSection = Section{
	Title: "🔧 SYSTEMD",
	Items: []Item{
		{Type: TypeTip, Value: "systemd — менеджер служб и целей загрузки.\nЮнит = описание службы (.service), сокета, таймера, target…"},

		{Type: TypeHeader, Value: "👀 Статус и логи"},
		{Type: TypeCmd, Value: "systemctl status service", Desc: "статус службы (active/failed, последние логи)"},
		{Type: TypeCmd, Value: "sudo journalctl -u service -f", Desc: "логи в реальном времени"},
		{Type: TypeCmd, Value: "sudo journalctl -u service -n 100", Desc: "последние 100 строк логов"},
		{Type: TypeCmd, Value: "sudo journalctl -xe", Desc: "свежие ошибки по системе"},
		{Type: TypeCmd, Value: "systemctl --failed", Desc: "список упавших служб"},

		{Type: TypeHeader, Value: "▶️ Управление службой"},
		{Type: TypeCmd, Value: "sudo systemctl start service", Desc: "запустить"},
		{Type: TypeCmd, Value: "sudo systemctl stop service", Desc: "остановить"},
		{Type: TypeCmd, Value: "sudo systemctl restart service", Desc: "перезапустить"},
		{Type: TypeCmd, Value: "sudo systemctl reload service", Desc: "перечитать конфиг без полного рестарта (если умеет)"},
		{Type: TypeCmd, Value: "sudo systemctl enable service", Desc: "включить автозапуск при загрузке"},
		{Type: TypeCmd, Value: "sudo systemctl disable service", Desc: "убрать из автозапуска"},
		{Type: TypeCmd, Value: "sudo systemctl enable --now service", Desc: "включить автозапуск и сразу запустить"},
		{Type: TypeCmd, Value: "systemctl is-enabled service", Desc: "enabled / disabled / static"},
		{Type: TypeCmd, Value: "systemctl is-active service", Desc: "active / inactive / failed"},

		{Type: TypeHeader, Value: "📄 Юниты и конфиги"},
		{Type: TypeCmd, Value: "systemctl cat service", Desc: "показать unit-файл (как systemd его видит)"},
		{Type: TypeCmd, Value: "systemctl list-units --type=service", Desc: "список служб"},
		{Type: TypeCmd, Value: "systemctl list-unit-files --type=service", Desc: "все unit-файлы и их enable-статус"},
		{Type: TypeCmd, Value: "sudo systemctl daemon-reload", Desc: "после правки unit-файла — обязательно"},
		{Type: TypeTip, Value: "Свои юниты: /etc/systemd/system/name.service\nПосле правки: daemon-reload → restart."},

		{Type: TypeHeader, Value: "🎯 Targets (уровни загрузки)"},
		{Type: TypeKey, Key: "multi-user.target", Desc: "обычный сервер без GUI"},
		{Type: TypeKey, Key: "graphical.target", Desc: "с графикой"},
		{Type: TypeKey, Key: "rescue.target", Desc: "режим восстановления"},
		{Type: TypeCmd, Value: "systemctl get-default", Desc: "текущий default target"},
		{Type: TypeCmd, Value: "sudo systemctl set-default multi-user.target", Desc: "сделать серверный режим по умолчанию"},

		{Type: TypeHeader, Value: "💡 Советы"},
		{Type: TypeTip, Value: "Не работает служба: status → journalctl -u … -xe → правишь конфиг → restart."},
		{Type: TypeWarn, Value: "disable не останавливает текущий процесс — только убирает автозапуск. Для «выключить сейчас» нужен stop."},
	},
}
