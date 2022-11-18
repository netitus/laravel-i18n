Laravel Internationalization Helper
---

This little tool will help you easily extract new translations 
from your Laravel (Blade) applications, by comparing Git changes.

# Usage

The only required parameter is the translation JSON in a locale of your choice.<br>
The default commit start point is `master`.

Run the command within your project's Git root:
```
/path/to/bin/laravel-i18n -t example.json 
```
Or between commits:
```
/path/to/bin/laravel-i18n -c hash123 -toc hsah321 -t example.json 
```

Please refer to the internal documentation for all options:
```
laravel-i18n --help
```