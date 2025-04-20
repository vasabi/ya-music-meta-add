# ya-music-meta-add для версий Ямузыки из Microsoft Store. Тестировалось на 4.54
Что делает скрипт:
  - Обрабатывает локально сохраненный кэш Ямузыки в формате mp3
  - Добавляет файлам названия и теги
  - Копирует в целевую директорию по пути Исполнитель -> Альбом -> Трек
Как пользоваться:
1 - Установить Ямузыку из Microsoft Store.
2 - Скачать кэш музыки локально
2 - Скачать файл ya-music-meta-add.exe
3 - В директории с файлом создать конфиг .cobra.yaml.
4 - Добавить в .cobra.yaml пути до базы данных, папки с mp3 файлами и выходную директорию
    Обычно это что-то вроде:
    music_path: "C:\Users\user_name\AppData\Local\Packages\A025C540.Yandex.Music_vfvw9svesycw6\LocalState\Music\35ee6bdbbe9142a9fa9dce0686e1d19e"
    db_path: "C:\Users\user_name\AppData\Local\Packages\A025C540.Yandex.Music_vfvw9svesycw6\LocalState\musicdb_35ee6bdbbe9142a9fa9dce0686e1d19e.sqlite"
    output_path: "F:\Музыка\YA_Music_Downloaded"
5 - Открыть базу данных любым удобным браузером, например DB Brouser for SQLite
6 - Изменить Journal Mod на off (В DB Brouser for SQLite это вкладка Edit Pragmas)
7 - Открыть cmd или PowerShell с правами администратора
8 - Перейти в директорию со скаченным скриптом (cd E:\Programs\yama)
9 - Запустить скрипт командой .\ya-music-meta-add.exe metadata start
10 - Дождаться выполнения. Скрипт автоматически создаст все необходимые директории и пропустит уже скопированные и обработанные файлы.
