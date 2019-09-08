package facade

import (
	"bytes"
	"io"

	"github.com/spf13/viper"
	"github.com/spiegel-im-spiegel/books-data/api"
	"github.com/spiegel-im-spiegel/books-data/api/pa"
	"github.com/spiegel-im-spiegel/books-data/ecode"
	"github.com/spiegel-im-spiegel/books-data/entity"
	"github.com/spiegel-im-spiegel/errs"
)

//paapiParams is parameters for PA-API
type paapiParams struct {
	marketplace  string
	associateTag string
	accessKey    string
	secretKey    string
}

func getPaapiParams() (*paapiParams, error) {
	marketplace := viper.GetString("marketplace")
	if len(marketplace) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidAPIParameter, "marketplace is empty")
	}
	associateTag := viper.GetString("associate-tag")
	if len(associateTag) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidAPIParameter, "associate-tag is empty")
	}
	accessKey := viper.GetString("access-key")
	if len(accessKey) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidAPIParameter, "access-key is empty")
	}
	secretKey := viper.GetString("secret-key")
	if len(secretKey) == 0 {
		return nil, errs.Wrap(ecode.ErrInvalidAPIParameter, "secret-key is empty")
	}
	return &paapiParams{marketplace: marketplace, associateTag: associateTag, accessKey: accessKey, secretKey: secretKey}, nil
}

func createPAAPI(p *paapiParams, isbnFlag bool) api.API {
	return pa.New(
		pa.WithMarketplace(p.marketplace),
		pa.WithAssociateTag(p.associateTag),
		pa.WithAccessKey(p.accessKey),
		pa.WithSecretKey(p.secretKey),
		pa.WithEnableISBN(isbnFlag),
	)
}

func searchPAAPI(id string, p *paapiParams, isbnFlag, rawFlag bool) (io.Reader, error) {
	if rawFlag {
		return createPAAPI(p, isbnFlag).LookupRawData(id)
	}
	book, err := createPAAPI(p, isbnFlag).LookupBook(id)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithParam("id", id))
	}
	b, err := book.Format(tmpltPath)
	if err != nil {
		return nil, errs.Wrap(err, "", errs.WithParam("id", id))
	}
	return bytes.NewReader(b), nil
}

func findPAAPI(id string, p *paapiParams, isbnFlag bool) (*entity.Book, error) {
	return createPAAPI(p, isbnFlag).LookupBook(id)
}

/* Copyright 2019 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
