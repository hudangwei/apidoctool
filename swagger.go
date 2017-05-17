package main

type APIService struct {
	ServiceInterfaces []ServiceInterface
}

func NewAPIService(api *API) *APIService {
	var l []ServiceInterface
	for url, paths := range api.Paths {
		serviceInterface := NewServiceInterface()
		if paths.Get != nil {
			serviceInterface.loadData(url, "Get", paths.Get, api.Definitions)
		}
		if paths.Post != nil {
			serviceInterface.loadData(url, "Post", paths.Post, api.Definitions)
		}
		if paths.Put != nil {
			serviceInterface.loadData(url, "Put", paths.Put, api.Definitions)
		}
		if paths.Patch != nil {
			serviceInterface.loadData(url, "Patch", paths.Patch, api.Definitions)
		}
		if paths.Delete != nil {
			serviceInterface.loadData(url, "Delete", paths.Delete, api.Definitions)
		}
		if paths.Head != nil {
			serviceInterface.loadData(url, "Head", paths.Head, api.Definitions)
		}

		l = append(l, *serviceInterface)
	}
	return &APIService{
		ServiceInterfaces: l,
	}
}

type ServiceInterface struct {
	OperationID string
	Description string
	Method      string
	URI         string
	Params      []Parameter
	Body        []Parameter
	Definitions map[string]Definition
}

func NewServiceInterface() *ServiceInterface {
	return &ServiceInterface{
		Definitions: map[string]Definition{},
	}
}

func (s *ServiceInterface) loadData(uri, method string, p *Operation, d map[string]Definition) {
	s.OperationID = p.OperationID
	s.Description = p.Description
	s.Method = method
	s.URI = uri
	if len(p.Parameters) > 0 {
		for _, parameter := range p.Parameters {
			s.Params = append(s.Params, parameter)
		}
	}

	for _, resp := range p.Responses {
		if ref, ok := resp.Schema["myRef"]; ok {

			for k, v := range d[ref.(string)].Properties {
				if v.MyRef != "" {
					v.Type = d[v.MyRef].Type + " <" + d[v.MyRef].Title + ">"

					s.Definitions[d[v.MyRef].Title] = d[v.MyRef]
				}
				if v.Type == "array" {
					if _, ok := v.Items["myRef"]; ok {
						v.Type = "array <" + d[v.Items["myRef"]].Title + ">"

						s.Definitions[d[v.Items["myRef"]].Title] = d[v.Items["myRef"]]

					}
				}
				param := Parameter{
					Name:        k,
					Type:        v.Type,
					Description: v.Description,
				}
				s.Body = append(s.Body, param)
			}
		}
		if typ, ok := resp.Schema["type"]; ok {

			if typ.(string) == "array" {
				if items, ok := resp.Schema["items"]; ok {
					for k, v := range d[items.(map[string]interface{})["myRef"].(string)].Properties {
						if v.MyRef != "" {
							v.Type = d[v.MyRef].Type + " <" + d[v.MyRef].Title + ">"

							s.Definitions[d[v.MyRef].Title] = d[v.MyRef]
						}
						if v.Type == "array" {
							if _, ok := v.Items["myRef"]; ok {
								v.Type = "array <" + d[v.Items["myRef"]].Title + ">"

								s.Definitions[d[v.Items["myRef"]].Title] = d[v.Items["myRef"]]

							}
						}
						param := Parameter{
							Name:        k,
							Type:        v.Type,
							Description: v.Description,
						}
						s.Body = append(s.Body, param)
					}
				}
			}
		}
	}
}
