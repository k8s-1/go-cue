import (
	"github.com/k8s-1/app/best"
)

#objects: best.#Def & {
	#vals: {
		{
			name: "Charlie Cartwright"
			age:  999
		}

	}
}
yaml.MarshalStream([for _, o in #objects {o}])
