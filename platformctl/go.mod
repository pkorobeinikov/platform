module platformctl

go 1.17

require (
	github.com/carlmjohnson/requests v0.22.1
	github.com/drone/envsubst v1.0.3
	github.com/joho/godotenv v1.4.0
	github.com/mitchellh/go-homedir v1.0.0
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8
	github.com/pkorobeinikov/platform/platform-lib v0.0.0-00010101000000-000000000000
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.8.1
)

replace github.com/pkorobeinikov/platform/platform-lib => ../platform-lib

require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/mitchellh/mapstructure v1.4.1 // indirect
	github.com/pelletier/go-toml v1.9.3 // indirect
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f // indirect
	golang.org/x/sys v0.0.0-20210616045830-e2b7044e8c71 // indirect
	golang.org/x/text v0.3.6 // indirect
	gopkg.in/ini.v1 v1.62.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)
