package data

var RegexSection = Section{
	Title: "🧩 РЕГУЛЯРНЫЕ ВЫРАЖЕНИЯ",
	Items: []Item{
		{Type: TypeTip, Value: "Regex — шаблон «найди вот такой текст».\nРаботает с: grep -E, sed -E, awk.\nПиши шаблон в одинарных кавычках '...' — иначе bash сломает $ и *."},
		{Type: TypeTip, Value: "Проверка:\ngrep -Eo 'шаблон' файл\n-E = расширенные regex\n-o = показать ТОЛЬКО совпадение (удобно учиться)"},

		{Type: TypeHeader, Value: "🧱 Буквы конструктора"},
		{Type: TypeKey, Key: "^", Desc: "начало строки: ^ERROR = строка начинается с ERROR"},
		{Type: TypeKey, Key: "$", Desc: "конец строки: done$ = строка заканчивается на done"},
		{Type: TypeKey, Key: ".", Desc: "ЛЮБОЙ один символ (a.b найдёт a1b, axb…)"},
		{Type: TypeKey, Key: "\\.", Desc: "настоящая точка. Для IP/доменов всегда \\."},
		{Type: TypeKey, Key: "*", Desc: "0 или много раз предыдущего: a* = пусто, a, aa, aaa…"},
		{Type: TypeKey, Key: "+", Desc: "1 или много раз: [0-9]+ = хотя бы одна цифра"},
		{Type: TypeKey, Key: "?", Desc: "0 или 1 раз: https? = http или https"},
		{Type: TypeKey, Key: "{n}", Desc: "ровно n раз: [0-9]{4} = ровно 4 цифры"},
		{Type: TypeKey, Key: "{n,m}", Desc: "от n до m: [0-9]{2,5} = от 2 до 5 цифр"},
		{Type: TypeKey, Key: "a|b", Desc: "ИЛИ: error|fail = error или fail"},
		{Type: TypeKey, Key: "()", Desc: "скобки = группа: (ru|com)"},
		{Type: TypeKey, Key: "[abc]", Desc: "один из символов: a или b или c"},
		{Type: TypeKey, Key: "[0-9]", Desc: "любая цифра 0–9"},
		{Type: TypeKey, Key: "[A-Za-z]", Desc: "любая латинская буква"},
		{Type: TypeKey, Key: "[^0-9]", Desc: "всё КРОМЕ цифр"},
		{Type: TypeKey, Key: "[[:space:]]", Desc: "пробел или таб"},

		{Type: TypeHeader, Value: "📧 Email"},
		{Type: TypeCmd, Value: "grep -Eo '[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}' file.txt", Desc: "найти email"},
		{Type: TypeTip, Value: "Пример: user.name+tag@mail.ru\n\n[A-Za-z0-9._%+-]+   → имя ящика: user.name+tag\n@                   → собака\n[A-Za-z0-9.-]+      → домен: mail\n\\.[A-Za-z]{2,}      → зона: .ru  (.com .io …)\n\nПочему \\. ? Точка без \\ = «любой символ».\nПочему {2,} ? Зона минимум 2 буквы (.r не бывает)."},

		{Type: TypeHeader, Value: "📱 Телефон (РФ)"},
		{Type: TypeCmd, Value: "grep -Eo '\\+7[0-9]{10}|8[0-9]{10}' file.txt", Desc: "слитно: +79991234567 или 89991234567"},
		{Type: TypeTip, Value: "Разбор \\+7[0-9]{10}|8[0-9]{10}\n\n\\+7        → плюс и семёрка (\\+ = настоящий +)\n[0-9]{10}  → ровно 10 цифр после\n|          → ИЛИ\n8[0-9]{10} → восьмёрка + 10 цифр\n\nИтого: +7__________  или  8__________"},
		{Type: TypeCmd, Value: "grep -Eo '\\+7[ ()-]*[0-9]{3}[ ()-]*[0-9]{3}[ ()-]*[0-9]{2}[ ()-]*[0-9]{2}' file.txt", Desc: "с скобками/пробелами: +7 (999) 123-45-67"},
		{Type: TypeTip, Value: "Разбор «красивого» телефона:\n\n\\+7              → +7\n[ ()-]*          → пробелы/скобки/дефисы (сколько угодно)\n[0-9]{3}         → код: 999\n[ ()-]*[0-9]{3}  → 123\n[ ()-]*[0-9]{2}  → 45\n[ ()-]*[0-9]{2}  → 67"},

		{Type: TypeHeader, Value: "🌐 IP-адрес (IPv4)"},
		{Type: TypeCmd, Value: "grep -Eo '([0-9]{1,3}\\.){3}[0-9]{1,3}' file.txt", Desc: "типа 192.168.0.1"},
		{Type: TypeTip, Value: "Разбор ([0-9]{1,3}\\.){3}[0-9]{1,3}\n\n[0-9]{1,3}  → кусок из 1–3 цифр (192)\n\\.           → точка между кусками\n(...){3}     → такой кусок+точка ПОВТОРИТЬ 3 раза\n               → 192.  168.  0.\n[0-9]{1,3}  → последний кусок без точки: 1\n\nПолучается: A.B.C.D\nГрубо: может поймать и 999.999.999.999 — для логов обычно ок."},

		{Type: TypeHeader, Value: "🔗 URL"},
		{Type: TypeCmd, Value: "grep -Eo 'https?://[^[:space:]]+' file.txt", Desc: "http:// или https:// до пробела"},
		{Type: TypeTip, Value: "Разбор https?://[^[:space:]]+\n\nhttp        → буквы http\ns?          → s можно пропустить → http или https\n://         → буквально ://\n[^[:space:]]+ → любые символы, КРОМЕ пробела/таба\n               (всё до конца URL)"},
		{Type: TypeCmd, Value: "grep -Eo 'https://[^[:space:]]+' access.log", Desc: "только https"},

		{Type: TypeHeader, Value: "📅 Дата и время"},
		{Type: TypeCmd, Value: "grep -Eo '[0-9]{4}-[0-9]{2}-[0-9]{2}' file.txt", Desc: "2026-07-14"},
		{Type: TypeTip, Value: "Разбор [0-9]{4}-[0-9]{2}-[0-9]{2}\n\n[0-9]{4}  → год: 2026\n-         → дефис\n[0-9]{2}  → месяц: 07\n-         → дефис\n[0-9]{2}  → день: 14"},
		{Type: TypeCmd, Value: "grep -Eo '[0-9]{2}\\.[0-9]{2}\\.[0-9]{4}' file.txt", Desc: "14.07.2026"},
		{Type: TypeTip, Value: "Разбор [0-9]{2}\\.[0-9]{2}\\.[0-9]{4}\n\nдень  .  месяц  .  год\n14    .  07     .  2026\n\\. нужен — иначе . = любой символ"},
		{Type: TypeCmd, Value: "grep -Eo '[0-9]{2}:[0-9]{2}(:[0-9]{2})?' file.txt", Desc: "10:31 или 10:31:45"},
		{Type: TypeTip, Value: "Разбор [0-9]{2}:[0-9]{2}(:[0-9]{2})?\n\n[0-9]{2}:[0-9]{2}  → часы:минуты\n(:[0-9]{2})?       → секунды необязательны (? = 0 или 1 раз)"},

		{Type: TypeHeader, Value: "🆔 UUID"},
		{Type: TypeCmd, Value: "grep -Eio '[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}' file.txt", Desc: "8-4-4-4-12 hex"},
		{Type: TypeTip, Value: "Разбор: xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx\n\n[0-9a-f]{8}  → 8 hex-символов\n-             → дефис\n[0-9a-f]{4}  → 4 hex\n…и так 5 блоков: 8-4-4-4-12\n\n-i в grep = без регистра (A-F тоже)."},

		{Type: TypeHeader, Value: "📡 MAC-адрес"},
		{Type: TypeCmd, Value: "grep -Eio '([0-9a-f]{2}:){5}[0-9a-f]{2}' file.txt", Desc: "aa:bb:cc:dd:ee:ff"},
		{Type: TypeTip, Value: "Разбор ([0-9a-f]{2}:){5}[0-9a-f]{2}\n\n[0-9a-f]{2}:  → два hex + двоеточие (aa:)\n(...){5}      → повторить 5 раз: aa:bb:cc:dd:ee:\n[0-9a-f]{2}   → последний байт без двоеточия: ff\n\nКак IP: группа+разделитель N раз, потом хвост."},

		{Type: TypeHeader, Value: "🔏 Хэш (MD5 / SHA)"},
		{Type: TypeCmd, Value: "grep -Eo '[0-9a-f]{32}|[0-9a-f]{40}|[0-9a-f]{64}' file.txt", Desc: "32=MD5, 40=SHA1, 64=SHA256"},
		{Type: TypeTip, Value: "Разбор [0-9a-f]{32}|[0-9a-f]{40}|[0-9a-f]{64}\n\nпросто: ровно 32 ИЛИ 40 ИЛИ 64 hex-символа подряд.\nДлина выдаёт тип хэша."},

		{Type: TypeHeader, Value: "🚪 Порт"},
		{Type: TypeCmd, Value: "grep -Eo ':[0-9]{2,5}' file.txt", Desc: ":80 :8080 :65535"},
		{Type: TypeTip, Value: "Разбор :[0-9]{2,5}\n\n:           → двоеточие перед портом\n[0-9]{2,5}  → от 2 до 5 цифр (80 … 65535)"},

		{Type: TypeHeader, Value: "🌍 Домен"},
		{Type: TypeCmd, Value: "grep -Eo '[A-Za-z0-9.-]+\\.(ru|com|org|io|net|dev)' file.txt", Desc: "site.ru example.com"},
		{Type: TypeTip, Value: "Разбор [A-Za-z0-9.-]+\\.(ru|com|org|io|net|dev)\n\n[A-Za-z0-9.-]+  → имя: example, mail, sub.domain\n\\.               → точка перед зоной\n(ru|com|…)      → одна из зон из списка"},

		{Type: TypeHeader, Value: "📂 Путь в Linux"},
		{Type: TypeCmd, Value: "grep -Eo '/[A-Za-z0-9._/-]+' file.txt", Desc: "/var/log/nginx/error.log"},
		{Type: TypeTip, Value: "Разбор /[A-Za-z0-9._/-]+\n\n/                 → начинается со слэша\n[A-Za-z0-9._/-]+  → буквы, цифры, точка, _, /, -"},

		{Type: TypeHeader, Value: "🔐 Секреты в логах"},
		{Type: TypeCmd, Value: "grep -Ei 'password|passwd|secret|token|api[_-]?key' app.log", Desc: "подозрительные слова"},
		{Type: TypeTip, Value: "Разбор api[_-]?key\n\napi       → буквы api\n[_-]?     → _ или - или ничего\nkey       → key\n\nНайдёт: apikey, api_key, api-key"},
		{Type: TypeCmd, Value: "grep -Eo 'Bearer [A-Za-z0-9._-]+' app.log", Desc: "Bearer-токен"},
		{Type: TypeTip, Value: "Разбор Bearer [A-Za-z0-9._-]+\n\nBearer␠     → слово + пробел\n[A-Za-z0-9._-]+  → сам токен"},
		{Type: TypeCmd, Value: "grep -Eo 'AKIA[0-9A-Z]{16}' file.txt", Desc: "AWS Access Key Id"},
		{Type: TypeTip, Value: "Разбор AKIA[0-9A-Z]{16}\n\nAKIA           → типичный префикс AWS ключа\n[0-9A-Z]{16}   → ещё 16 символов"},
		{Type: TypeWarn, Value: "Нашёл секрет — сразу ротируй. В git такое не коммить."},

		{Type: TypeHeader, Value: "🧹 Комментарии и пустые строки"},
		{Type: TypeCmd, Value: "grep -E '^#' file.txt", Desc: "строка начинается с #"},
		{Type: TypeCmd, Value: "grep -E '^$' file.txt", Desc: "совсем пустая строка"},
		{Type: TypeCmd, Value: "grep -E '^[[:space:]]*$' file.txt", Desc: "пустая или только пробелы"},
		{Type: TypeTip, Value: "Разбор ^[[:space:]]*$\n\n^              → начало\n[[:space:]]*   → пробелы/табы 0 или много\n$              → конец\n= между началом и концом ничего, кроме пробелов"},
		{Type: TypeCmd, Value: "sed -E '/^[[:space:]]*#/d;/^[[:space:]]*$/d' file.txt", Desc: "убрать комментарии и пустые (на экран)"},
		{Type: TypeCmd, Value: "sed -i -E '/^[[:space:]]*#/d;/^[[:space:]]*$/d' file.txt", Desc: "то же — сохранить в файл"},
		{Type: TypeTip, Value: "Разбор sed '/шаблон/d'\n\n/шаблон/  → найти такие строки\nd         → delete (удалить)\n;         → вторая команда следом\n-i        → записать в файл"},

		{Type: TypeHeader, Value: "🐳 Docker / образы"},
		{Type: TypeCmd, Value: "grep -Eo '[A-Za-z0-9._/-]+:[A-Za-z0-9._-]+' file.txt", Desc: "nginx:1.25  repo/app:latest"},
		{Type: TypeTip, Value: "Разбор [A-Za-z0-9._/-]+:[A-Za-z0-9._-]+\n\n[A-Za-z0-9._/-]+  → имя образа (можно с /)\n:                  → двоеточие перед тегом\n[A-Za-z0-9._-]+   → тег: 1.25 / latest"},
		{Type: TypeCmd, Value: "grep -E 'EXITED|OOM|killed|error' docker.log", Desc: "типичные сбои в логах"},

		{Type: TypeHeader, Value: "✍️ Замаскировать (sed)"},
		{Type: TypeCmd, Value: "sed -E 's/[0-9]{4}-[0-9]{2}-[0-9]{2}/**DATE**/g' file.txt", Desc: "даты → **DATE**"},
		{Type: TypeCmd, Value: "sed -E 's/[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\\.[A-Za-z]{2,}/**EMAIL**/g' file.txt", Desc: "email → **EMAIL**"},
		{Type: TypeCmd, Value: "sed -E 's/\\+7[0-9]{10}/**PHONE**/g' file.txt", Desc: "телефон → **PHONE**"},
		{Type: TypeTip, Value: "Формула sed: s/ЧТО_ИЩЕМ/НА_ЧТО_МЕНЯЕМ/g\n\ns = substitute (заменить)\ng = global (все вхождения в строке)\nСначала без -i — посмотреть. Потом -i — сохранить."},

		{Type: TypeHeader, Value: "💡 Как читать любой шаблон"},
		{Type: TypeTip, Value: "1) Разбей на куски слева направо\n2) Скобки (...){n} = «этот кусок n раз»\n3) Точка почти всегда должна быть \\.\n4) Проверь: grep -Eo 'шаблон' файл"},
		{Type: TypeWarn, Value: "Забыл \\. перед зоной/IP — шаблон начнёт ловить мусор."},
	},
}
