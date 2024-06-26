// Copyright (c) 2021 Terminus, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package models

import (
	"github.com/erda-project/erda-infra/providers/i18n"
	"github.com/erda-project/erda/bundle"
)

type Service struct {
	db       *DBClient
	bundle   *bundle.Bundle
	i18nTran i18n.Translator
	lang     i18n.LanguageCodes
}

func NewService(dbClient *DBClient, bundle *bundle.Bundle, i18nTran i18n.Translator, lang i18n.LanguageCodes) *Service {
	s := Service{}
	s.db = dbClient
	s.bundle = bundle
	s.i18nTran = i18nTran
	s.lang = lang
	return &s
}

func (svc *Service) I18n(key string) string {
	return svc.i18nTran.Text(svc.lang, key)
}
