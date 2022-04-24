package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetWebsites = map[string]struct{}{
	/* Other countries */

	"https://online.sberbank.kz/CSAFront/index.do": {},

	/* Russia */

	// Propaganda
	"https://lenta.ru":                           {},
	"https://tass.ru":                            {},
	"https://tass.ru/userApi/getNewsFeed":        {},
	"https://vz.ru":                              {},
	"https://www.gazeta.ru":                      {},
	"https://api.riafan.ru":                      {},
	"https://pikabu.ru":                          {},
	"https://api.pikabu.ru":                      {},
	"https://i.kommersant.ru":                    {},
	"https://omk-job.ru":                         {},
	"https://www.rambler.ru":                     {},
	"https://www.rambler.ru/api/v4/personalized": {},
	// 	"https://id.rambler.ru/login-20/login": {}, // TODO fix small read buffer
	"https://mail.rambler.ru":                            {},
	"https://news.rambler.ru":                            {},
	"https://peroxide.rambler.ru/v1/comments/clusters/":  {},
	"https://rabota.rambler.ru":                          {},
	"https://rcm.rambler.ru":                             {},
	"https://mail.ru":                                    {},
	"https://www.astrobl.ru":                             {},
	"https://www2.astrobl.ru":                            {},
	"https://adm.astrobl.ru":                             {},
	"https://transport.stavregion.ru":                    {},
	"https://rg.ru":                                      {},
	"http://allpravda.info":                              {},
	"https://antifashist.com":                            {}, // https://antifashist.online
	"https://antimaydan.info":                            {},
	"https://argumenti.ru":                               {},
	"https://aurora.network":                             {},
	"https://balrad.ru":                                  {},
	"https://buzina.org":                                 {},
	"https://www.c-inform.info":                          {},
	"https://crimea-news.com":                            {},
	"https://crimeapress.info":                           {},
	"https://comitet.su":                                 {},
	"https://compromat.group":                            {},
	"https://compromat.group/engine/ajax/checkViews.php": {},
	"https://cont.ws":                                    {},
	"https://dan-news.info":                              {},
	"https://dni.ru":                                     {},
	"https://dnr24.com":                                  {},
	"https://dnr24.su":                                   {},
	"https://dnr-life.ru":                                {},
	"https://dnr-pravda.ru":                              {},
	"https://donbasstoday.ru":                            {},
	"https://dontimes.ru":                                {},
	"https://dosie.su":                                   {}, // cf
	"https://eadaily.com/ru/dossier":                     {},
	"https://eadaily.com/ru/rss/index.xml":               {},
	"https://evening-crimea.com":                         {},
	"https://expert.ru":                                  {},
	"https://fine-news.ru":                               {},
	"https://finobzor.ru":                                {},
	"https://forpostsevastopol.ru":                       {},
	"https://front-novorossii.ru":                        {},
	"http://globalwarnews.ru":                            {},
	"https://gorlovka-news.su":                           {},
	"https://gtrklnr.com":                                {},
	"https://inforeactor.ru":                             {},
	"https://infox.sg":                                   {},
	"https://inlugansk.ru":                               {},
	"https://jpgazeta.ru":                                {},
	"https://kianews24.ru":                               {},
	"http://k-politika.ru":                               {},
	"https://kuban24.tv":                                 {},
	"https://life.ru":                                    {},
	"https://api.corr.life/public/regions/lookup":        {},
	"https://lug-info.com":                               {},
	"https://lugansk-online.su":                          {},
	"https://www.m24.ru":                                 {},
	"https://cross.m24.ru/covid/frontend/web/site/rus":   {},
	"https://mediarepost.ru":                             {},
	"https://www.metronews.ru":                           {},
	"https://www.readmetro.com":                          {},
	"https://www.mk-kalm.ru":                             {},
	"https://mkset.ru":                                   {},
	"http://www.moscow-post.su":                          {},
	"https://www.nakanune.ru":                            {},
	"https://nation-news.ru":                             {},
	"https://newizv.ru":                                  {},
	"https://newizv.ru/api/v1/matters":                   {},
	"http://news24today.info":                            {},
	"https://newsland.com":                               {},
	"https://nvo.ng.ru":                                  {},
	"https://nnews.nnov.ru":                              {},
	"https://newdaynews.ru":                              {},
	"https://www.newc.info":                              {},
	// 	"https://odnarodyna.org":                             {}, // TODO small read buffer
	"https://novorossiia.ru":              {}, // cf
	"https://novorossiia.info":            {},
	"http://novorossy.ru":                 {},
	"https://novosti.icu":                 {},
	"http://novosti333.ru":                {},
	"https://nt1941.su":                   {},
	"https://pandoraopen.ru":              {},
	"https://plainnews.ru":                {},
	"https://pravdoryb.info":              {},
	"https://polit.info":                  {},
	"https://politcentr.ru":               {},
	"https://politexpert.net":             {},
	"https://politikus.ru":                {},
	"https://politpuzzle.ru":              {},
	"https://politros.com":                {},
	"https://rf-smi.ru":                   {},
	"https://www.ritmeurasia.org":         {},
	"https://rodina.news":                 {}, // cf
	"https://rossaprimavera.ru":           {},
	"https://www.company.rt.ru":           {},
	"https://rusdnepr.ru":                 {},
	"https://rusnext.ru":                  {}, // cf
	"https://rusnext.ru/newslistapi.json": {}, // cf
	"http://ru-an.info":                   {},
	"https://rueconomics.ru":              {},
	"https://rusonline.org":               {},
	"http://ruspravda.info":               {},
	"https://sevnews.info":                {},
	"https://slovodel.com":                {},
	"https://smi2.ru":                     {},
	"https://polls.smi2.ru/body/1/poll":   {},
	"https://sovross.ru":                  {},
	"http://stringer-news.com":            {},
	"https://svpressa.ru":                 {},
	"https://taurica.net":                 {},
	"https://tehnowar.ru":                 {},
	"https://time-news.net":               {}, // cf
	"https://trmzk.ru":                    {},
	"https://u-f.ru":                      {},
	"https://ugyalta.com":                 {}, // cf
	"https://universe-tss.su":             {},
	"https://utro.ru":                     {},
	"https://www.vedomosti.ru":            {},
	"https://vesti-k.ru":                  {},
	"https://vm.ru":                       {},
	"https://voskhodinfo.su":              {},
	// 	"https://webkamerton.ru":              {}, // small read buffer
	"http://webnovosti.info":  {},
	"https://wpristav.ru":     {}, // cf
	"https://x-true.info":     {}, // cf
	"http://xvesti.ru":        {},
	"https://zanamipravda.ru": {},
	"https://zavtra.ru":       {},

	// https://www.state.gov/russias-pillars-of-disinformation-and-propaganda-report/
	"http://www.geopolitika.ru":                 {},
	"https://katehon.com":                       {},
	"https://novorosinform.org":                 {},
	"https://www.pravda.ru":                     {},
	"https://www.pravda.ru/ajaxed/toolbartabs/": {},
	"https://russia-insider.com/ru":             {},
	"https://southfront.org":                    {},
	"https://www.strategic-culture.org":         {},
	"https://therussophile.org":                 {},
	"https://tsargrad.tv":                       {},
	"https://www.voltairenet.org":               {}, // cf

	// Business corporations
	"http://www.crimearw.ru":                                      {},
	"https://www.wildberries.ru":                                  {},
	"https://api-stories.wildberries.ru/api/v1/articles/previews": {},
	"https://chat.wildberries.ru/v1/unauth/messages":              {},
	"https://guru.wildberries.ru":                                 {},
	"https://suggestions.wildberries.ru/api/v2/hint":              {},
	"https://wbxcatalog-ru.wildberries.ru/tops_tshirts/catalog":   {},
	"https://wbxsearch.wildberries.ru/exactmatch/v2/common":       {},
	"https://www.ozon.ru":                                         {},
	"https://www.ozon.ru/api/composer-api.bx/page/json/v2":        {},
	"https://brandlab.ozon.ru":                                    {},
	"https://docs.ozon.ru/main":                                   {},
	"https://job.ozon.ru":                                         {},
	"https://seller.ozon.ru":                                      {},
	"https://www.dns-shop.ru":                                     {},
	"https://advego.com":                                          {},
	"https://freelance.ru":                                        {},
	"https://new.ingosinvest.ru/Registration/Info":                {},
	"https://pif.ingosinvest.ru":                                  {},
	"https://ok.ru":                                               {},
	"http://www.yemelya.ru":                                       {},
	"https://vcs.rostec.ru":                                       {},
	"https://vks3.rostec.ru":                                      {},
	"https://kontur.ru":                                           {},
	"http://app.rzd.ru":                                           {},
	"https://team.rzd.ru":                                         {},
	"https://rzdint.ru":                                           {},
	"https://www.rzdlog.com":                                      {},

	//The state
	"https://mil.ru":                                   {},
	"https://www.nalog.gov.ru":                         {}, // should be called without params
	"https://pfr.gov.ru":                               {},
	"https://gosuslugi41.ru":                           {},
	"https://otpravka.pochta.ru":                       {},
	"https://passport.pochta.ru/oauth2/authorize":      {},
	"https://tariff.pochta.ru":                         {},
	"https://zakaznoe.pochta.ru":                       {},
	"http://crimea-post.ru":                            {},
	"https://ca.vks.rosguard.gov.ru":                   {},
	"https://sozd.duma.gov.ru":                         {},
	"https://vsednr.ru":                                {},
	"https://ecis.admoblkaluga.ru":                     {},
	"https://entry.admoblkaluga.ru/Web/Login":          {},
	"https://mimz.admoblkaluga.ru/Menu/Page/1":         {},
	"https://mimz.admoblkaluga.ru/GzwSP/ContractsJson": {},
	"https://pre.admoblkaluga.ru/main/":                {},
	"https://sadko30.admoblkaluga.ru/account/login":    {},
	"https://soc.admoblkaluga.ru":                      {},
	"https://spo.admoblkaluga.ru/security/":            {},
	"https://tender.admoblkaluga.ru":                   {},
	"https://dnmchs.ru":                                {},
	"http://www.fsb.ru":                                {},
	"https://govdnr.ru":                                {},
	"https://mer.govdnr.ru":                            {},
	"https://mid-dnr.su/ru/":                           {},
	"https://agmo.mosreg.ru":                           {},
	"https://br.mosreg.ru/login/auth/":                 {},
	"https://dk.mosreg.ru":                             {},
	"https://easuz.mosreg.ru":                          {},
	"https://gasu.mosreg.ru":                           {},
	"https://invest.mosreg.ru":                         {},
	"https://pass.mosreg.ru/login":                     {},
	"https://login.school.mosreg.ru":                   {},
	"https://web.mail.mosreg.ru/SOGo/":                 {},
	"https://market.mosreg.ru":                         {},
	"https://rgis.mosreg.ru/v3/":                       {},
	"https://sso.mosreg.ru/signin.jsp":                 {},
	"https://support.mosreg.ru":                        {},
	"https://uslugi.mosreg.ru":                         {},
	"https://vmeste.mosreg.ru/login":                   {},
	"https://welcome.mosreg.ru":                        {},
	"http://wow.mosreg.ru/admin/":                      {},
	"https://edu-mosreg.ru":                            {},
	"https://mzdnr.ru":                                 {},
	"https://pravdnr.ru":                               {},
	"https://supcourt-dpr.su":                          {},
	"https://ulgov.ru":                                 {},

	// Embassies
	// Do not duplicate the IPs: https://github.com/erkexzcx/stoppropaganda/pull/110#issuecomment-1059960305
	"https://montreal.mid.ru": {},
	"https://belarus.mid.ru":  {},

	// Others
	"https://www.glonass-iac.ru": {},

	// Electronic signature services, certificate authorities, www domain names
	"https://etp.kartoteka.ru/index.html": {},
	"http://mcspro.ru":                    {},

	// Cinemas
	"https://karofilm.ru":   {},
	"https://www.mirage.ru": {},

	// Belgorod
	"https://bel.ru": {},
	"https://bel.ru/api/v1/platform/main_page":                 {},
	"https://belgorod.chibbis.ru":                              {},
	"https://belgorod.igooods.ru":                              {},
	"https://belgorod.igooods.ru/api/v8/cities/identify_by_ip": {},
	"https://belgorod.farfor.ru":                               {},
	"https://belsbyt.ru":                                       {},
	"https://lk.belsbyt.ru/fiz/login":                          {},
	"https://www.go31.ru":                                      {},
	"https://www.go31.ru/api3/auth/check-login":                {},
	"https://perevozki31.ru":                                   {},

	// Others
	"http://217.12.104.100": {},

	// Various websites by ip
	"https://213.59.254.8":   {},
	"https://82.151.111.187": {},

	/* BELARUS */

	// by gov
	"https://president.gov.by/ru":   {},
	"https://mfa.gov.by":            {},
	"https://russia.mfa.gov.by/ru/": {},

	// by banks
	"https://belarusbank.by": {},

	// by business
	"https://www.rw.by": {},

	// by media
	"https://www.belta.by":       {},
	"https://www.belarus.by/ru/": {},
	"https://ont.by":             {}, // should be called without params
	"https://www.belnovosti.by":  {},

	/* Syria */
	"https://syrianfinance.gov.sy": {},

	// Still operating in Russia
	// https://www.yalerussianbusinessretreat.com/
	"https://transport.auchan.ru/core/framework/login.cfm":          {},
	"https://burgerkingrus.ru":                                      {},
	"https://burgerkingrus.ru/api-web-front/api/v3/restaurant/list": {},
	"https://mic.burgerking.ru/mifs/user/login.jsp":                 {},
	"https://sd.burgerking.ru/HEAT/SaaSExternalSessionRenew.aspx":   {},
	"https://www.decathlon.ru":                                      {}, // should be called without params
	"https://blog.decathlon.ru":                                     {},
	"https://delonghi.ru":                                           {},
	"https://ecco.ru":                                               {},
	"https://www.jd.ru":                                             {},
	"https://www.jd.ru/leproxy/api/carts/cart/owner-id":             {},
	"https://mega.ru/offers/":                                       {}, // Ingka
	"https://modshairrussia.ru":                                     {},
	"https://mymodshair.ru":                                         {},
	"https://natr.ru":                                               {},
	"https://papajohns.ru/moscow":                                   {},
	"https://api.papajohns.ru/slider/list":                          {},
	"https://subway.ru":                                             {},
	"https://www.yves-rocher.ru":                                    {},

	// https://t.me/itarmyofukraine2022/230
	"https://www.mid.ru": {},

	// https://t.me/itarmyofukraine2022/215
	"https://www.chechnya.online": {},
	"https://vaynahavia.com":      {},
	"https://www.minfinchr.ru":    {},

	// https://t.me/itarmyofukraine2022/287
	"https://www.amediateka.ru": {},
	"https://www.okko.tv":       {},
	"https://wink.ru":           {},

	// https://t.me/itarmyofukraine2022/289
	"https://stream.1tv.ru/api/playlist/1tvch-v1_as_array.json": {},

	// https://t.me/itarmyofukraine2022/302
	"https://rov.aero": {},
}
