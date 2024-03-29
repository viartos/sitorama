package core

import (
	"context"

	"github.com/art-sitedesign/sitorama/app/core/services"
	"github.com/art-sitedesign/sitorama/app/utils"
)

// Init произведет инициализацию приложения, запустив нужные контейнеры
func (c *Core) Init(ctx context.Context) error {
	// создание подсети приложения
	networkID, err := c.initNetwork(ctx)
	if err != nil {
		return err
	}

	// создание контейнера-роутера запросов по исполняющим контейнерам
	containerID, err := c.initRouter(ctx)
	if err != nil {
		return err
	}

	// подключение контейнера-роутера к подсети приложения
	err = c.docker.ConnectNetwork(ctx, networkID, containerID, []string{utils.RouterName})
	if err != nil {
		return err
	}

	// запуск созданного контейнера
	err = c.docker.StartContainer(ctx, containerID)
	if err != nil {
		return err
	}

	return nil
}

func (c *Core) initNetwork(ctx context.Context) (string, error) {
	netID := ""
	res, err := c.docker.FindNetwork(ctx)
	if err != nil {
		return netID, err
	}

	if res == nil {
		netID, err = c.docker.CreateNetwork(ctx)
		if err != nil {
			return netID, err
		}

		return netID, nil
	}

	netID = res.ID
	return netID, nil
}

func (c *Core) initRouter(ctx context.Context) (string, error) {
	router := services.NewRouter(c.docker)

	res, err := router.Find(ctx)
	if err != nil {
		return "", err
	}

	if res == nil {
		cID, err := router.Create(ctx)
		if err != nil {
			return "", err
		}

		return cID, nil
	}

	return res.ID, nil
}
