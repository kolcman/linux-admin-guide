package data

var CronSection = Section{
	Title: "⏰ CRON — ПЛАНИРОВЩИК",
	Items: []Item{
		{Type: TypeTip, Value: "Cron запускает команды и скрипты по расписанию.\nСвои задачи — crontab. Общие/пакетные — файлы в /etc/cron.d/ (каталог .d: много маленьких конфигов вместо одного большого)."},

		{Type: TypeHeader, Value: "🛠️ crontab"},
		{Type: TypeCmd, Value: "crontab -l", Desc: "список задач текущего пользователя"},
		{Type: TypeCmd, Value: "crontab -e", Desc: "редактировать (первый раз спросит редактор)"},
		{Type: TypeCmd, Value: "ls /etc/cron.d/", Desc: "системные сценарии (по файлу на задачу или группу)"},
		{Type: TypeCmd, Value: "journalctl -u cron", Desc: "логи службы cron"},

		{Type: TypeHeader, Value: "📋 Формат строки"},
		{Type: TypeTip, Value: "мин час день месяц день_недели команда\n* = каждый. Задача идёт от того пользователя, кто её создал."},
		{Type: TypeKey, Key: "минута", Desc: "0–59"},
		{Type: TypeKey, Key: "час", Desc: "0–23"},
		{Type: TypeKey, Key: "день месяца", Desc: "1–31"},
		{Type: TypeKey, Key: "месяц", Desc: "1–12 или jan,feb,mar…"},
		{Type: TypeKey, Key: "день недели", Desc: "0–6 (0 или 7 = воскресенье)"},
		{Type: TypeKey, Key: "*", Desc: "каждое значение поля"},
		{Type: TypeKey, Key: "*/N", Desc: "каждые N единиц: */10 в минутах = каждые 10 мин"},
		{Type: TypeKey, Key: "1-5", Desc: "диапазон: пн–пт"},

		{Type: TypeHeader, Value: "⏱️ Примеры расписания"},
		{Type: TypeKey, Key: "* * * * *", Desc: "каждую минуту"},
		{Type: TypeKey, Key: "*/10 * * * *", Desc: "каждые 10 минут"},
		{Type: TypeKey, Key: "03 05 * * *", Desc: "ежедневно в 05:03"},
		{Type: TypeKey, Key: "0 22 * * 1-5", Desc: "в 22:00 по будням"},
		{Type: TypeKey, Key: "59 23 31 dec *", Desc: "31 декабря в 23:59"},
		{Type: TypeKey, Key: "0 3 * * *", Desc: "каждый день в 03:00"},
		{Type: TypeKey, Key: "* * * * * date >> /tmp/date", Desc: "строка в crontab: каждую минуту дописать дату"},
		{Type: TypeCmd, Value: "cat /tmp/date", Desc: "проверить, что задача отработала"},

		{Type: TypeHeader, Value: "@ сокращения (не везде)"},
		{Type: TypeKey, Key: "@hourly", Desc: "раз в час = 0 * * * *"},
		{Type: TypeKey, Key: "@daily", Desc: "полночь = 0 0 * * *"},
		{Type: TypeKey, Key: "@weekly", Desc: "полночь воскресенья = 0 0 * * 0"},
		{Type: TypeKey, Key: "@monthly", Desc: "полночь 1-го числа = 0 0 1 * *"},
		{Type: TypeKey, Key: "@yearly", Desc: "1 января 00:00 = 0 0 1 1 *"},
		{Type: TypeWarn, Value: "@hourly/@daily… — нестандарт. На части UNIX могут не работать. Надёжнее пять полей."},

		{Type: TypeHeader, Value: "💡 Советы"},
		{Type: TypeTip, Value: "Вывод задачи пиши в файл: date >> /tmp/date  или  cmd >> /var/log/job.log 2>&1"},
		{Type: TypeWarn, Value: "Cron работает, только пока включена ОС. Выключенная VM задачу пропустит."},
	},
}
