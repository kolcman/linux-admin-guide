package data

var BackupSection = Section{
	Title: "📦 БЭКАПЫ",
	Items: []Item{
		// TAR.GZ
		{Type: TypeCmd, Value: "tar -czvf archive.tar.gz folder/", Desc: "Создать TAR.GZ архив"},
		{Type: TypeCmd, Value: "tar -xzvf archive.tar.gz", Desc: "Распаковать TAR.GZ архив"},

		// ZIP
		{Type: TypeCmd, Value: "zip -r archive.zip folder/", Desc: "Создать ZIP архив"},
		{Type: TypeCmd, Value: "unzip archive.zip", Desc: "Распаковать ZIP архив"},

		// 7Z
		{Type: TypeCmd, Value: "7z a archive.7z folder/", Desc: "Создать 7Z архив"},
		{Type: TypeCmd, Value: "7z x archive.7z", Desc: "Распаковать 7Z архив"},

		// Бэкапы
		{Type: TypeCmd, Value: "rsync -avh src/ dst/", Desc: "Синхронизация папок"},
		{Type: TypeCmd, Value: "crontab -e", Desc: "Редактор планировщика задач"},
	},
}
