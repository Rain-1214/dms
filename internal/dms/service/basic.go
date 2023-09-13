package service

import (
	"context"

	v1 "github.com/actiontech/dms/api/dms/service/v1"
	"github.com/actiontech/dms/internal/dms/biz"
)

func (d *DMSService) GetBasicInfo(ctx context.Context) (reply *v1.GetBasicInfoReply, err error) {
	basic, err := d.BasicUsecase.GetBasicInfo(ctx)
	if nil != err {
		return nil, err
	}

	ret := &v1.BasicInfo{
		LogoUrl: basic.LogoUrl,
		Title:   basic.Title,
	}

	for _, item := range basic.Components {
		ret.Components = append(ret.Components, v1.ComponentNameWithVersion{
			Name:    item.Name,
			Version: item.Version,
		})
	}

	return &v1.GetBasicInfoReply{
		Payload: struct {
			BasicInfo *v1.BasicInfo `json:"basic_info"`
		}{ret},
	}, nil
}

func (d *DMSService) GetStaticLogo(ctx context.Context) (*v1.GetStaticLogoReply, string, error) {
	basicConfig, contentType, err := d.BasicUsecase.GetStaticLogo(ctx)
	if nil != err {
		return nil, contentType, err
	}

	return &v1.GetStaticLogoReply{
		File: basicConfig.Logo,
	}, contentType, nil
}

func (d *DMSService) Personalization(ctx context.Context, req *v1.PersonalizationReq) error {
	params := &biz.BasicConfigParams{
		Title: req.Title,
		File:  req.File,
	}

	return d.BasicUsecase.Personalisation(ctx, params)
}