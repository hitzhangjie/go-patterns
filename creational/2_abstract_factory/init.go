package abstract_factory

// Style defines style of products
type Style int

const (
	Style1 = Style(iota)
	Style2
)

var factories = map[Style]IFactory{}

func init() {
	Register(Style1Factory{})
	Register(Style2Factory{})
}

func Register(f IFactory) {
	factories[f.Style()] = f
}

func Get(style Style) IFactory {
	return factories[style]
}
