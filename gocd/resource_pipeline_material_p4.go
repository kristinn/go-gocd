package gocd

import "errors"

func (mp4 MaterialAttributesP4) equal(mp42i interface{}) (bool, error) {
	var ok bool
	mp42, ok := mp42i.(*MaterialAttributesP4)
	if !ok {
		return false, errors.New("can only compare with same material type")
	}

	namesEqual := mp4.Name == mp42.Name
	portEqual := mp4.Port == mp42.Port
	destEqual := mp4.Destination == mp42.Destination

	return namesEqual && portEqual && destEqual, nil
}

// UnmarshallInterface from a JSON string to a MaterialAttributesP4 struct
func unmarshallMaterialAttributesP4(mp4 *MaterialAttributesP4, i map[string]interface{}) {
	for key, value := range i {
		if value == nil {
			continue
		}
		switch key {
		case "name":
			mp4.Name = value.(string)
		case "port":
			mp4.Port = value.(string)
		case "use_tickets":
			mp4.UseTickets = value.(bool)
		case "view":
			mp4.View = value.(string)
		case "username":
			mp4.Username = value.(string)
		case "password":
			mp4.Password = value.(string)
		case "encrypted_password":
			mp4.EncryptedPassword = value.(string)
		case "destination":
			mp4.Destination = value.(string)
		case "filter":
			mp4.Filter = unmarshallMaterialFilter(value)
		case "invert_filter":
			mp4.InvertFilter = value.(bool)
		case "auto_update":
			mp4.AutoUpdate = value.(bool)
		}
	}
}
