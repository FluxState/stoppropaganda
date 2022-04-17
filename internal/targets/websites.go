package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetWebsites = map[string]struct{}{
	/* Other countries */

	"https://bukimevieningi.lt":                    {},
	"https://musutv.lt":                            {},
	"https://www.viada.lt":                         {},
	"https://online.sberbank.kz/CSAFront/index.do": {},

	/* Russia */

	// Propaganda
	"https://tass.ru":                            {},
	"https://tass.ru/userApi/getNewsFeed":        {},
	"https://www.rzd.ru":                         {},
	"https://rzdlog.ru":                          {},
	"https://vz.ru":                              {},
	"https://api.riafan.ru":                      {},
	"https://pikabu.ru":                          {},
	"https://api.pikabu.ru":                      {},
	"https://i.kommersant.ru":                    {},
	"https://omk-job.ru":                         {},
	"https://www.rambler.ru":                     {},
	"https://www.rambler.ru/api/v4/personalized": {},
	// 	"https://id.rambler.ru/login-20/login": {}, // TODO fix small read buffer
	"https://mail.rambler.ru":                           {},
	"https://news.rambler.ru":                           {},
	"https://peroxide.rambler.ru/v1/comments/clusters/": {},
	"https://rabota.rambler.ru":                         {},
	"https://rcm.rambler.ru":                            {},
	"https://mail.ru":                                   {},
	"https://www.astrobl.ru":                            {},
	"https://www2.astrobl.ru":                           {},
	"https://adm.astrobl.ru":                            {},
	"https://mosreg.ru":                                 {},
	"https://easuz.mosreg.ru":                           {},
	"https://login.school.mosreg.ru":                    {},
	"https://uslugi.mosreg.ru":                          {},
	"https://edu-mosreg.ru":                             {},
	"https://ulgov.ru":                                  {},
	"https://transport.cheladmin.ru":                    {},
	"https://stavregion.ru":                             {},
	"http://dostup.stavregion.ru/dsreda/login":          {},
	"http://lk.stavregion.ru":                           {},
	"https://torgi.stavregion.ru/application/":          {},
	"https://transport.stavregion.ru":                   {},
	"https://rg.ru":                                     {},
	"https://balrad.ru":                                 {},
	"https://crimea-news.com":                           {},
	"https://eadaily.com/ru/dossier":                    {},
	"https://eadaily.com/ru/rss/index.xml":              {},
	"https://lug-info.com":                              {},
	"https://kuban24.tv":                                {},
	"https://newdaynews.ru":                             {},

	// https://www.state.gov/russias-pillars-of-disinformation-and-propaganda-report/
	"http://www.geopolitika.ru":                 {},
	"https://www.globalresearch.ca":             {},
	"https://katehon.com":                       {},
	"https://news-front.info":                   {},
	"https://novorosinform.org":                 {},
	"https://www.pravda.ru":                     {},
	"https://www.pravda.ru/ajaxed/toolbartabs/": {},
	"https://russia-insider.com/ru":             {},
	"https://southfront.org":                    {},
	"https://www.strategic-culture.org":         {},
	"https://therussophile.org":                 {},
	"https://tsargrad.tv":                       {},
	"https://www.voltairenet.org":               {},

	// Business corporations
	"https://www.evraz.com/ru/":                            {},
	"https://nlmk.com/ru/":                                 {},
	"https://www.wildberries.ru":                           {},
	"https://www.ozon.ru":                                  {},
	"https://www.ozon.ru/api/composer-api.bx/page/json/v2": {},
	"https://www.dns-shop.ru":                              {},
	"https://aliexpress.ru":                                {},
	"https://advego.com":                                   {},
	"https://freelance.ru":                                 {},
	"https://www.turbotext.ru":                             {},
	"https://www.vtbcapital-pr.ru":                         {},
	"https://www.ingosinvest.ru":                           {},
	"https://ok.ru":                                        {},
	"http://5.61.23.11":                                    {},
	"http://217.20.155.13":                                 {},
	"http://217.20.147.1":                                  {},
	"http://www.yemelya.ru":                                {},
	"https://scloud.rostec.ru/login":                       {},
	"https://vcs.rostec.ru":                                {},
	"https://vks3.rostec.ru":                               {},
	"https://kontur.ru":                                    {},
	"https://help.kontur.ru/ke":                            {},

	// Banks
	"https://sber.ru": {},
	"https://api.developer.sber.ru/product/SberbankID": {},
	"https://promo.sber.ru/sberserenity":               {},
	"https://sberprime.sber.ru":                        {},
	"https://www.sberbank.com/ru":                      {},
	"https://www.sberbank.ru":                          {},
	"https://online.sberbank.ru":                       {},
	"https://api.sberbank.ru/prod/tokens/v2":           {},
	"https://api.sberbank.ru/prod/tokens/v2/oauth":     {},
	"https://api.sberbank.ru/prod/tokens/v2/oidc":      {},
	"https://online.sberbank.ru/CSAFront/index.do":     {},
	"https://www.vtb.ru":                               {},
	"https://cl.vtb.ru":                                {},
	"https://ipoteka.vtb.ru/ipoteka/1":                 {},
	// 	"https://online.vtb.ru/login": {}, // TODO small read buffer
	"https://school.vtb.ru":                                          {},
	"https://siteapi.vtb.ru/api/sitepages/component-types":           {},
	"https://www.moex.com":                                           {},
	"https://iss.moex.com/iss/reference/":                            {},
	"https://messaging.moex.com/init":                                {},
	"https://passport.moex.com/login":                                {},
	"http://www.fsb.ru":                                              {},
	"https://scr.online.sberbank.ru/api/fl/idgib-w-3ds":              {},
	"https://acs1.sbrf.ru":                                           {},
	"https://acs2.sbrf.ru":                                           {},
	"https://acs3.sbrf.ru":                                           {},
	"https://acs4.sbrf.ru":                                           {},
	"https://acs5.sbrf.ru":                                           {},
	"https://acs6.sbrf.ru":                                           {},
	"https://acs7.sbrf.ru":                                           {},
	"https://acs8.sbrf.ru":                                           {},
	"https://enter.unicredit.ru/v2/cgi/bsi.dll?T=RT_2Auth.BF":        {},
	"https://online.gpb.ru/login":                                    {},
	"https://www.rshb.ru":                                            {},
	"https://online.rshb.ru/b1/ib6/wf2/retail/ib/loginretaildefault": {},
	"https://online.mkb.ru":                                          {},
	"https://id.tinkoff.ru/auth/step":                                {},
	"https://131.ru":                                                 {},
	"https://developer.131.ru":                                       {},
	"https://online.131.ru":                                          {},
	"https://online.131.ru/json/auth/login":                          {},
	"https://www.akbars.ru":                                          {},
	"https://bankok.akbars.ru/AkbarsOnlineAuth/SavedLoginInfo":       {},
	"https://invest.akbars.ru":                                       {},
	"https://invest.akbars.ru/api/o":                                 {},
	"https://online.akbars.ru":                                       {},
	"https://altbank.ru/v2/s/bsi.dll":                                {},
	"https://apkbank.ru/ru/msk/":                                     {},
	"https://www.aresbank.ru":                                        {},
	"https://www.baikalcredobank.ru/irk/":                            {},
	"https://www.banki.ru":                                           {},
	"https://www.banki.ru/ng/api/v1.0/public/banners":                {},
	"https://faktura.ru":                                             {},
	"https://bank.faktura.ru":                                        {},
	"https://business.faktura.ru/f2b/":                               {},
	"https://elf.faktura.ru/elf/app/":                                {},

	//The state
	"https://mil.ru":                              {},
	"https://www.nalog.gov.ru":                    {},
	"https://pfr.gov.ru":                          {},
	"https://rkn.gov.ru":                          {},
	"https://gosuslugi41.ru":                      {},
	"https://otpravka.pochta.ru":                  {},
	"https://passport.pochta.ru/oauth2/authorize": {},
	"https://tariff.pochta.ru":                    {},
	"https://sozd.duma.gov.ru":                    {},
	"https://wagnera.ru":                          {},

	// Embassies
	// Do not duplicate the IPs: https://github.com/erkexzcx/stoppropaganda/pull/110#issuecomment-1059960305
	"https://montreal.mid.ru": {},
	"https://belarus.mid.ru":  {},

	// Others
	"https://proverki-gov.ru":    {},
	"https://www.glonass-iac.ru": {},
	"https://market.dme.ru":      {},
	"https://www.tickets.dme.ru": {},

	// Payments
	"http://185.170.2.7":         {},
	"http://185.170.3.7":         {},
	"https://ds1.mirconnect.ru":  {},
	"https://ds2.mirconnect.ru":  {},
	"https://wats.mirconnect.ru": {},
	"https://passport.etpgpb.ru": {},
	"https://bki-okb.ru":         {},

	// Exchanges connected to russian banks
	"https://bitokk.biz":      {},
	"https://www.netex24.net": {},
	"https://flashobmen.com":  {},
	"https://yoomoney.ru":     {},
	"https://yookassa.ru":     {},

	// Electronic signature services, certificate authorities, www domain names
	"https://www.roseltorg.ru":            {},
	"https://178fz.roseltorg.ru":          {},
	"https://agro.roseltorg.ru":           {},
	"https://atom2.roseltorg.ru":          {},
	"https://bg.roseltorg.ru/auth/email":  {},
	"https://com.roseltorg.ru":            {},
	"https://docs.roseltorg.ru":           {},
	"https://etp.roseltorg.ru":            {},
	"https://fkr.roseltorg.ru":            {},
	"https://fkr2.roseltorg.ru":           {},
	"https://kim-atom.roseltorg.ru":       {},
	"https://kim.roseltorg.ru":            {},
	"https://kim-irao.roseltorg.ru":       {},
	"https://lk.roseltorg.ru/auth":        {},
	"https://orders.roseltorg.ru":         {},
	"https://rosgeo.roseltorg.ru":         {},
	"https://rosseti.roseltorg.ru":        {},
	"https://rt.roseltorg.ru":             {},
	"https://rushydro.roseltorg.ru":       {},
	"https://msp.roseltorg.ru":            {},
	"https://vtb.roseltorg.ru":            {},
	"https://rosinvoice.ru/auth":          {},
	"https://kb.kubankredit.ru/login":     {},
	"https://structure.mil.ru":            {},
	"https://parus-s.ru":                  {},
	"https://www.icentr.ru":               {},
	"https://etp.kartoteka.ru/index.html": {},
	"https://www.24ecp.ru":                {},
	"https://squaretrade.ru":              {},
	"https://www.cit-ufa.ru":              {},
	"https://api.cit-ufa.ru":              {},
	"https://icvibor.ru":                  {},
	"https://1c.icvibor.ru":               {},
	"http://mcspro.ru":                    {},
	"https://iecp.ru/ep/ep-verification":  {},
	"https://e-trust.gosuslugi.ru":        {},
	"https://gu.spb.ru":                   {},
	"https://iecp.ru/ep/uc-list":          {},

	// Oil and gas trading companies
	"https://www.tektorg.ru":                                      {},
	"https://44.tektorg.ru/authentication/login":                  {},
	"https://fin.tektorg.ru/bg":                                   {}, // static
	"https://fin.tektorg.ru/api/v2/reference/service_types":       {}, // api pass
	"https://zakupki.tektorg.ru":                                  {},
	"https://aaa-srt-cs.lukoil.com/logon/LogonPoint/tmindex.html": {},
	"https://ar.lukoil.com/login":                                 {},
	"https://b2b.sibur.ru/pages_new_en/faces/index.jspx":          {},
	"https://onlinecontract.ru":                                   {},

	// Food delivery services
	"https://sbermarket.ru": {},
	"https://chibbis.ru":    {},
	"https://samokat.ru":    {},

	// Cinemas
	"https://karofilm.ru":   {},
	"https://www.mirage.ru": {},

	// Transport and logistics
	"https://a-groupp.com":                     {},
	"https://atec-logistic.ru":                 {},
	"https://forums.ati.su/Forum/Default.aspx": {},
	"https://help.ati.su":                      {},
	"https://news.ati.su/index/":               {},
	"https://s1.ati.su/api/6/envelope":         {},
	"https://trace.ati.su":                     {},
	"https://trucks.ati.su":                    {},
	"https://trucks.ati.su/webapi/public/v1.0/loads/dictionaries/get_deals_dicts": {},
	"https://userdata.ati.su/api/iewv_r":                                          {},
	"https://zen.ati.su":                                                          {},
	"https://zen.ati.su/api/v1/top/tags":                                          {},
	"https://avtobus-raspisanie.ru":                                               {},
	"https://avtobus.flydex.ru":                                                   {},
	"https://avtobus.flydex.ru/ajax/suggest/buses":                                {},
	"https://avtoperevozki-rostov.ru":                                             {},
	"https://azlog.ru":                                                            {},
	"https://bilet.aero":                                                          {},
	"https://avia.bilet.aero/api/guide/airports/MOW":                              {},
	"https://busnovoyas.ru":                                                       {},
	"https://www.dellin.ru":                                                       {},
	"https://www.dellin.ru/api/calculation/v1/products":                           {},
	"https://kogda.by":                                                            {},
	"https://www.tutu.ru":                                                         {},
	"https://bus.tutu.ru/api/v1/geo/suggest":                                      {},
	"https://id.tutu.ru/sdk_api/session":                                          {},
	"https://story-proxy.tutu.ru/main_page_posts/":                                {},

	// Belgorod
	"https://bel.ru": {},
	"https://bel.ru/api/v1/platform/main_page":                 {},
	"https://belgorod.igooods.ru/api/v8/cities/identify_by_ip": {},
	"https://belgorod.farfor.ru":                               {},
	"https://belsbyt.ru":                                       {},
	"https://lk.belsbyt.ru/fiz/login":                          {},
	"https://www.go31.ru":                                      {},
	"https://www.go31.ru/api3/auth/check-login":                {},
	"https://perevozki31.ru":                                   {},

	// Others
	"http://217.12.104.100":  {},
	"https://toxic-repos.ru": {},

	// Various websites by ip
	"https://91.213.144.193": {},
	"https://91.213.144.237": {},
	"https://212.24.38.190":  {},
	"https://78.47.115.99":   {},
	"https://194.190.37.226": {},
	"https://194.190.37.228": {},
	"https://213.59.254.8":   {},
	"https://185.194.34.142": {},
	"https://213.59.197.65":  {},
	"https://185.71.67.237":  {},
	"https://82.151.111.187": {},
	"https://91.239.5.75":    {},
	"https://92.53.98.191":   {},

	/* BELARUS */

	// by gov
	"https://president.gov.by/ru":   {},
	"https://www.kgb.by/ru":         {},
	"https://mfa.gov.by":            {},
	"https://russia.mfa.gov.by/ru/": {},

	// by banks
	"https://belarusbank.by": {},

	// by business
	"https://www.rw.by": {},

	// by media
	"https://www.belta.by":      {},
	"https://www.sb.by":         {},
	"https://www.belarus.by":    {},
	"https://ont.by":            {},
	"https://www.024.by":        {},
	"https://www.belnovosti.by": {},
	"https://radiomir.by":       {},
	"https://api.radiomir.by":   {},
	"https://grodnonews.by":     {},

	/* Syria */
	"https://syrianfinance.gov.sy": {},
	"http://185.216.132.201":       {},

	// Still operating in Russia
	// https://yale.box.com/s/11lqy1d3yn1kf9xa3r96k9sb6w5m4qea
	"https://www.albane.ru":    {},
	"https://burgerkingrus.ru": {},
	"https://burgerkingrus.ru/api-web-front/api/v3/restaurant/list": {},
	"https://delonghi.ru":                                         {},
	"https://mic.burgerking.ru/mifs/user/login.jsp":               {},
	"https://sd.burgerking.ru/HEAT/SaaSExternalSessionRenew.aspx": {},
	"https://www.decathlon.ru":                                    {},
	"https://ecco.ru":                                             {},
	"https://lladro.ru":                                           {},
	"https://mega.ru":                                             {}, // Ingka
	"https://modshairrussia.ru":                                   {},
	"https://mymodshair.ru":                                       {},
	"https://natr.ru":                                             {},
	"https://papajohns.ru/moscow":                                 {},
	"https://api.papajohns.ru/slider/list":                        {},
	"https://subway.ru":                                           {},
	"https://www.yves-rocher.ru":                                  {},

	// https://t.me/itarmyofukraine2022/230
	"http://www.ved.gov.ru":      {},
	"https://www.mid.ru":         {},
	"https://www.economy.gov.ru": {},

	// https://t.me/itarmyofukraine2022/215
	"https://www.chechnya.online": {},
	"https://vaynahavia.com":      {},
	"https://www.minfinchr.ru":    {},

	"https://www.company.rt.ru": {},

	// https://t.me/itarmyofukraine2022/243
	"https://fss.gov.ru":              {},
	"https://autodiscover.fss.ru/owa": {},
	"https://portal.fss.ru":           {},
	"https://data.fss.ru/open":        {},
	"http://docs.fss.ru":              {},
	"https://hosting.pfrf.ru":         {},
	"https://school.pfr.gov.ru":       {},
	"https://es-pfrf.ru/":             {},

	// https://t.me/itarmyofukraine2022/247
	"https://qiwi.com":          {},
	"https://checkout.qiwi.com": {},
	"https://my.qiwi.com":       {},
	"https://oplata.qiwi.com":   {},
	"https://p2p.qiwi.com":      {},

	// https://t.me/itarmyofukraine2022/248
	"https://digital.ac.gov.ru":       {},
	"https://digital.gov.ru":          {},
	"https://lk.digital.gov.ru":       {},
	"https://platform.digital.gov.ru": {},
	"https://reestr.digital.gov.ru":   {},
	"https://sc.digital.gov.ru/home":  {},
	"https://minenergo.gov.ru":        {},
	"https://minfin.gov.ru/ru/":       {},

	// https://t.me/itarmyofukraine2022/251
	"https://payanyway.ru/info/w/ru/public/welcome.htm": {},

	// https://t.me/itarmyofukraine2022/253
	"https://www.superjob.ru": {},
	"https://rabota.vk.com":   {},

	// https://t.me/itarmyofukraine2022/269
	"https://cdek.by/ru/":         {},
	"https://www.cdek.ru/graphql": {},

	// https://t.me/itarmyofukraine2022/273
	"https://www.doski.ru":                           {},
	"https://www.farpost.ru":                         {},
	"https://unibo.ru":                               {},
	"https://api.youla.io/api/v1/search/suggestions": {},

	// https://t.me/itarmyofukraine2022/275
	"https://www.crpt.ru":                  {},
	"https://ismp.crpt.ru":                 {},
	"https://kb.crpt.ru":                   {},
	"https://markirovka.crpt.ru/login-kep": {},
	"https://api.mdlp.crpt.ru":             {},
	"https://api.sb.mdlp.crpt.ru":          {},
	"https://nrz.api.sb.mdlp.crpt.ru":      {},
	"https://milk.crpt.ru":                 {},
	"https://suzgrid.crpt.ru":              {},

	// https://t.me/itarmyofukraine2022/287
	"https://www.amediateka.ru": {},
	"https://www.okko.tv":       {},
	"https://www.ontvtime.ru":   {},
	"https://wink.ru":           {},

	// https://t.me/itarmyofukraine2022/289
	"https://stream.1tv.ru/api/playlist/1tvch-v1_as_array.json": {},
	"https://mt.kino-teatr.ru":                                  {},

	// https://t.me/itarmyofukraine2022/290
	"https://api.kontur.ru/csi/support/v1/public/integrations/zakupki": {},
	"https://auth.kontur.ru":                       {},
	"https://ca.kontur.ru":                         {},
	"https://developer.kontur.ru":                  {},
	"https://elbank.kontur.ru/AccessControl/Login": {},
	"https://focus.kontur.ru":                      {},
	"https://focus.kontur.ru/api/lists/all":        {},
	"https://install.kontur.ru/kekep":              {},
	"https://pf.kontur.ru":                         {},
	"https://zakupki.kontur.ru":                    {},
	"https://www.b-kontur.ru":                      {},
	"https://www.kontur-extern.ru":                 {},

	// https://t.me/itarmyofukraine2022/298
	"https://www.tvigle.ru":          {},
	"https://admin.tvigle.ru":        {},
	"https://cloud.tvigle.ru":        {},
	"https://downloader.tvigle.ru":   {},
	"http://gr.tvigle.ru":            {},
	"https://vendor.tvigle.ru/login": {},
}
