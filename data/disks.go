package data

var DisksSection = Section{
	Title: "💾 ДИСКИ — МОНТИРОВАНИЕ И АВТОЗАГРУЗКА",
	Items: []Item{
		{Type: TypeHeader, Value: "🔑 Ключи lsblk:"},
		{Type: TypeKey, Key: "-f", Desc: "Показать файловые системы (ext4, xfs и т.д.)"},
		{Type: TypeKey, Key: "-m", Desc: "Показать права доступа и владельца"},
		{Type: TypeKey, Key: "-o NAME,SIZE,TYPE,MOUNTPOINT", Desc: "Вывести только выбранные колонки"},

		{Type: TypeHeader, Value: "🛠️ Базовые команды:"},
		{Type: TypeCmd, Value: "lsblk -f", Desc: "Список дисков, UUID и точек монтирования"},
		{Type: TypeCmd, Value: "sudo fdisk -l", Desc: "Детальная таблица разделов (MBR/GPT)"},
		{Type: TypeCmd, Value: "df -h", Desc: "Свободное место на смонтированных дисках"},

		{Type: TypeHeader, Value: "📝 Сценарий подключения нового диска:"},
		{Type: TypeCmd, Value: "sudo parted /dev/sdb mklabel gpt", Desc: "1. Создать таблицу разделов GPT (для дисков >2TB)"},
		{Type: TypeCmd, Value: "sudo parted /dev/sdb mkpart primary ext4 0% 100%", Desc: "2. Создать один раздел на весь диск"},
		{Type: TypeCmd, Value: "sudo mkfs.ext4 /dev/sdb1", Desc: "3. Отформатировать раздел в ext4"},
		{Type: TypeCmd, Value: "sudo blkid /dev/sdb1", Desc: "4. Скопировать UUID раздела"},
		{Type: TypeCmd, Value: "sudo mkdir -p /mnt/data", Desc: "5. Создать точку монтирования"},

		{Type: TypeHeader, Value: "⚙️ Настройка автомонтирования (/etc/fstab):"},
		{Type: TypeCmd, Value: "sudo nano /etc/fstab", Desc: "6. Открыть файл конфигурации"},
		{Type: TypeCmd, Value: "UUID=ваш-uuid /mnt/data ext4 defaults 0 2", Desc: "7. Добавить строку (вместо ваш-uuid вставить UUID из шага 4)"},
		{Type: TypeCmd, Value: "sudo mount -a", Desc: "8. Проверить корректность fstab (монтирует всё)"},
	},
}
