package providers

var (
	None = ""
)

type ICPEngine interface {
	Name() string                            //name of engine
	SearchICP(domain string) (string, error) // main function to get icp
}
