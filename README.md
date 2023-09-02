# CS2 logaddress_add_http "Polyfill"

Counter-Strike 2 currently does not have the `logaddress_add_http` command available for use.

While I'm sure this will eventually be implemented, as HLTV is reliant on it, I have created this basic workaround.

Using the `-condebug` launch parameter, the server will write to a log file. By default at `C:\Program Files (x86)\Steam\steamapps\common\Counter-Strike Global Offensive\game\csgo\console.log`.

This will check for updates every 1/128th of a second. If there are any, it'll group them up and shoot them off at whichever URL you supply.

## Usage

Used in Powershell on Windows 10.

```
.\cs2_logaddress_add_http.exe -url "http://localhost:3000/receive" -file "C:\Program Files (x86)\Steam\steamapps\common\Counter-Strike Global Offensive\game\csgo\console.log"
```