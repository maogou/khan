package khan

import (
	"context"
	"maogou/khan/api/khan/v1/transform/contact"
	"maogou/khan/internal/pkg/log"
)

func (k *Khan) ContactList(ctx context.Context, req contact.ContactListRequest) (*contact.ContactListResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactListResponse{}).Post(contactList)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactList方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactListResponse), nil
}

func (k *Khan) ContactSearch(ctx context.Context, req contact.ContactSearchRequest) (*contact.ContactSearchResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactSearchResponse{}).Post(contactSearch)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSearch方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactSearchResponse), nil
}

func (k *Khan) ContactAdd(ctx context.Context, req contact.ContactAddRequest) (*contact.ContactAddResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactAddResponse{}).Post(contactAdd)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactAdd方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactAddResponse), nil
}

func (k *Khan) ContactDel(ctx context.Context, req contact.ContactDelRequest) (*contact.ContactDelResponse, error) {

	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactDelResponse{}).Post(contactDel)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactDel方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactDelResponse), nil
}

func (k *Khan) ContactInfo(ctx context.Context, req contact.ContactInfoRequest) (*contact.ContactInfoResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactInfoResponse{}).Post(contactInfo)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactInfo方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactInfoResponse), nil
}

func (k *Khan) ContactDetail(ctx context.Context, req contact.ContactDetailRequest) (*contact.ContactDetailResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactDetailResponse{}).Post(contactDetail)
	if err != nil {
		log.C(ctx).Error().Msg("调用contactDetail方法失败")
		return nil, err
	}
	return resp.Result().(*contact.ContactDetailResponse), nil
}

func (k *Khan) ContactSetOnlyChat(ctx context.Context, req contact.ContactSetOnlyChatRequest) (*contact.ContactSetOnlyChatResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactSetOnlyChatResponse{}).Post(contactSetOnlyChat)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSetOnlyChat方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactSetOnlyChatResponse), nil
}

func (k *Khan) ContactSetRemark(ctx context.Context, req contact.ContactSetRemarkRequest) (*contact.ContactSetRemarkResponse, error) {

	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactSetRemarkResponse{}).Post(contactSetRemark)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactSetRemark方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactSetRemarkResponse), nil
}

func (k *Khan) ContactUploadMobile(ctx context.Context, req contact.ContactUploadMobileRequest) (*contact.ContactUploadMobileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactUploadMobileResponse{}).Post(contactUploadMobile)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactUploadMobile方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactUploadMobileResponse), nil
}

func (k *Khan) ContactGetMobile(ctx context.Context, req contact.ContactGetMobileRequest) (*contact.ContactGetMobileResponse, error) {
	resp, err := k.client.R().SetBody(req).SetResult(&contact.ContactGetMobileResponse{}).Post(contactGetMobile)

	if err != nil {
		log.C(ctx).Error().Msg("调用contactGetMobile方法失败")
		return nil, err
	}

	return resp.Result().(*contact.ContactGetMobileResponse), nil
}
