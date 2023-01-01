package themes

import "github.com/pablodz/gnome-theme-manager/utils"

func GetLocalThemes() ([]string, error) {

	themes, err := utils.GetDirectories("/usr/share/themes")
	if err != nil {
		return nil, err
	}

	return themes, nil
}
