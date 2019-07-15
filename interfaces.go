package godatainterfaces

type IEnvCollection interface {
	GetEnvVariable(key string) string
	ListEnvVariables() []string
	ToMap() map[string]string
}

type IApplicationRuntime interface {
	GetEnvironmentVars() IEnvCollection
	GetAppProperties() IPropertyCollection
	GetClient(key string) interface{}
	GetLogger() ILogger
}

type IPropertyCollection interface {
	LockPropertyValues()
	GetValue(key string) string
	SetValue(key string, value string) IPropertyCollection
	MergeProperties(properties IPropertyCollection)
	ListProperties() []string
	ToMap() map[string]string
	Subscribe(key string) chan string
	Snapshot() IPropertyCollection
	Clone() IPropertyCollection
}

type ILogger interface {
	Log(string)
	LogWithLevel(level string, loginfo string)
}

type IResolveSelector interface {
	GetResolver(url string) IResolver
}

type IResolver interface {
	Exists() bool
	GetMetadata() map[string]string
	GetData() chan []byte
}

