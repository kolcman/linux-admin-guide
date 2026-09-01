package data

var BackupSection = Section{
	Title: "📦 БЭКАПЫ",
	Items: []Item{
		{Type: TypeTip, Value: "Бэкап — не «когда всё упало», а привычка.\nМинимум: уметь собрать архив нужных каталогов и синхронизировать через rsync."},

		{Type: TypeHeader, Value: "📦 TAR.GZ"},
		{Type: TypeKey, Key: "-c", Desc: "создать архив"},
		{Type: TypeKey, Key: "-x", Desc: "извлечь"},
		{Type: TypeKey, Key: "-z", Desc: "gzip-сжатие"},
		{Type: TypeKey, Key: "-v", Desc: "показывать файлы"},
		{Type: TypeKey, Key: "-f", Desc: "имя файла архива"},
		{Type: TypeKey, Key: "--exclude", Desc: "не класть в архив (можно несколько раз)"},
		{Type: TypeCmd, Value: "tar -czvf archive.tar.gz folder/", Desc: "создать TAR.GZ архив"},
		{Type: TypeCmd, Value: "tar -xzvf archive.tar.gz", Desc: "распаковать TAR.GZ"},
		{Type: TypeCmd, Value: "tar -czvf backup.tar.gz --exclude='node_modules' --exclude='.git' project/", Desc: "архив без мусора"},

		{Type: TypeHeader, Value: "📦 ZIP / 7Z"},
		{Type: TypeCmd, Value: "zip -r archive.zip folder/", Desc: "создать ZIP архив"},
		{Type: TypeCmd, Value: "unzip archive.zip", Desc: "распаковать ZIP"},
		{Type: TypeCmd, Value: "7z a archive.7z folder/", Desc: "создать 7Z архив"},
		{Type: TypeCmd, Value: "7z x archive.7z", Desc: "распаковать 7Z"},

		{Type: TypeHeader, Value: "🔄 rsync"},
		{Type: TypeKey, Key: "-a", Desc: "архивный режим: права, времена, рекурсия"},
		{Type: TypeKey, Key: "-v", Desc: "подробный вывод"},
		{Type: TypeKey, Key: "-h", Desc: "размеры по-человечески"},
		{Type: TypeKey, Key: "--delete", Desc: "удалить в dst то, чего нет в src (зеркало)"},
		{Type: TypeKey, Key: "--dry-run", Desc: "показать, что сделает, без копирования"},
		{Type: TypeCmd, Value: "rsync -avh src/ dst/", Desc: "синхронизация папок"},
		{Type: TypeCmd, Value: "rsync -avh --dry-run --delete src/ dst/", Desc: "показать зеркало, ничего не копировать"},
		{Type: TypeCmd, Value: "rsync -avh --delete src/ dst/", Desc: "зеркало: dst станет как src"},
		{Type: TypeCmd, Value: "rsync -avh -e ssh src/ user@host:/backup/", Desc: "бэкап на удалённый сервер по SSH"},
		{Type: TypeWarn, Value: "Слеш в конце важен: src/ — содержимое папки; src — сама папка целиком.\n--delete опасен: сначала --dry-run."},

		{Type: TypeHeader, Value: "⏰ По расписанию"},
		{Type: TypeTip, Value: "Планировщик: раздел ⏰ CRON. Для бэкапа достаточно одной строки в crontab."},
		{Type: TypeKey, Key: "0 3 * * * /usr/local/bin/backup.sh", Desc: "каждый день в 03:00"},
		{Type: TypeCmd, Value: "crontab -e", Desc: "добавить задачу"},
	},
}
