package data

var BootSection = Section{
	Title: "🚑 АВАРИЙНАЯ ЗАГРУЗКА",
	Items: []Item{
		{Type: TypeTip, Value: "Когда система не поднимается нормально — битый fstab, сломанный пароль root, упавший сервис на старте — нужны урезанные режимы загрузки."},
		{Type: TypeTip, Value: "Обычная загрузка кратко:\nUEFI/MBR → GRUB → ядро + initramfs → systemd → target (службы)."},

		{Type: TypeHeader, Value: "🔀 Режимы"},
		{Type: TypeKey, Key: "rescue.target", Desc: "режим восстановления: root, базовая система, часто с сетью"},
		{Type: TypeKey, Key: "emergency.target", Desc: "аварийный: минимум служб, корень часто только для чтения"},
		{Type: TypeKey, Key: "multi-user.target", Desc: "обычный сервер без GUI"},
		{Type: TypeTip, Value: "Имена в меню GRUB (recovery / rescue / emergency / single) зависят от дистрибутива — суть одна: урезанная загрузка для починки."},

		{Type: TypeHeader, Value: "🚪 Как зайти"},
		{Type: TypeTip, Value: "Через GRUB: Advanced / дополнительные параметры → recovery,\nили e на пункте ядра и добавь параметр systemd.unit=… (см. ниже), потом Ctrl+X / F10."},
		{Type: TypeCmd, Value: "systemd.unit=rescue.target", Desc: "параметр ядра в GRUB (правка e)"},
		{Type: TypeCmd, Value: "systemd.unit=emergency.target", Desc: "параметр ядра: emergency через GRUB"},

		{Type: TypeHeader, Value: "🛠️ Уже внутри системы"},
		{Type: TypeCmd, Value: "systemctl isolate rescue.target", Desc: "перейти в rescue с работающей системы"},
		{Type: TypeCmd, Value: "systemctl isolate emergency.target", Desc: "перейти в emergency"},
		{Type: TypeCmd, Value: "systemctl default", Desc: "вернуться к обычному multi-user/graphical"},
		{Type: TypeCmd, Value: "mount -o remount,rw /", Desc: "сделать корень RW в emergency (часто нужен)"},
		{Type: TypeCmd, Value: "systemctl reboot", Desc: "перезагрузка после правок"},

		{Type: TypeHeader, Value: "💡 Типовые поломки"},
		{Type: TypeTip, Value: "Битый fstab: загрузка зависает на mount.\nВойди в rescue/emergency → remount rw → nano /etc/fstab → mount -a → reboot."},
		{Type: TypeTip, Value: "Забыл пароль: recovery → remount rw → passwd username → reboot."},
		{Type: TypeWarn, Value: "На сервере с шифрованием диска или без пароля root вход в recovery может потребовать passphrase / root password.\n«Сеть недоступна» в recovery — нормально для части дистрибутивов."},
	},
}
