package pubchemclient

import "fmt"

func (c *PubChemClient) GetCompounds(domain, namespace, identifier, output string) (Compounds, error) {
	url := fmt.Sprintf("%s/%s/%s/%s/%s", commonURL, domain, namespace, identifier, output)
	compound, err := apiHandler[Compounds](c, url)
	if err != nil {
		return Compounds{}, err
	}
	return compound, nil
}
