package jingdong_union_go

import (
	"encoding/json"
	"errors"
)

type JdUnionOpenGoodsMaterialQueryTopLevel struct {
	JdUnionOpenGoodsMaterialQueryResponse JdUnionOpenGoodsMaterialQueryResponse `json:"jd_union_open_goods_material_query_responce"`
}

type JdUnionOpenGoodsMaterialQueryResponse struct {
	Result string `json:"queryResult"`
	Code   string `json:"code"`
}

func (app *App) JdUnionOpenGoodsMaterialQuery(params map[string]interface{}) (result *JdUnionOpenGoodsQueryResult, err error) {

	body, err := app.Request("jd.union.open.goods.material.query", map[string]interface{}{"goodsReq": params})

	resp := &JdUnionOpenGoodsMaterialQueryTopLevel{}
	if err != nil {
		// log.Println(string(body))
		return
	}
	if err = json.Unmarshal(body, resp); err != nil {
		return
	}
	if resp.JdUnionOpenGoodsMaterialQueryResponse.Result != "" {
		result = &JdUnionOpenGoodsQueryResult{}
		if err = json.Unmarshal([]byte(resp.JdUnionOpenGoodsMaterialQueryResponse.Result), result); err != nil {
			return
		}
	} else {
		err = errors.New("result is null")
	}
	return
}
