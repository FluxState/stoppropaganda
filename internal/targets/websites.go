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
	"https://novosti.icu": {},
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
	"https://www.evraz.com/ru/":                            {},
	"https://www.wildberries.ru":                           {},
	"https://www.ozon.ru":                                  {},
	"https://www.ozon.ru/api/composer-api.bx/page/json/v2": {},
	"https://www.dns-shop.ru":                              {},
	"https://advego.com":                                   {},
	"https://freelance.ru":                                 {},
	"https://www.turbotext.ru":                             {},
	"https://www.ingosinvest.ru":                           {},
	"https://ok.ru":                                        {},
	"http://5.61.23.11":                                    {},
	"http://217.20.155.13":                                 {},
	"http://217.20.147.1":                                  {},
	"http://www.yemelya.ru":                                {},
	"https://vcs.rostec.ru":                                {},
	"https://vks3.rostec.ru":                               {},
	"https://kontur.ru":                                    {},
	"https://www.rzd.ru":                                   {},
	"http://app.rzd.ru":                                    {},
	"https://team.rzd.ru":                                  {},
	"https://rzdint.ru":                                    {},
	"https://www.rzdlog.com":                               {},

	// Banks
	"https://sber.ru": {},
	"https://promo.sber.ru/sberserenity":               {},
	"https://sberprime.sber.ru":                        {},
	"https://www.sberbank.com/ru":                      {},
	"https://www.sberbank.ru":                          {},
	"https://api.sberbank.ru/prod/tokens/v2":           {},
	"https://api.sberbank.ru/prod/tokens/v2/oauth":     {},
	"https://api.sberbank.ru/prod/tokens/v2/oidc":      {},
	"https://online.sberbank.ru/CSAFront/index.do":     {},
	"https://cl.vtb.ru":                                {},
	"https://ipoteka.vtb.ru/ipoteka/1":                 {},
	// 	"https://online.vtb.ru/login": {}, // TODO small read buffer
	"https://school.vtb.ru":                                          {},
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
	"https://online.gpb.ru/login":                                    {},

	//The state
	"https://mil.ru":                              {},
	"https://www.nalog.gov.ru":                    {}, // should be called without params
	"https://pfr.gov.ru":                          {},
	"https://gosuslugi41.ru":                      {},
	"https://otpravka.pochta.ru":                  {},
	"https://passport.pochta.ru/oauth2/authorize": {},
	"https://tariff.pochta.ru":                    {},
	"https://sozd.duma.gov.ru":                    {},
	"https://vsednr.ru":                           {},
	"https://www.admoblkaluga.ru":                        {},
	"https://ecis.admoblkaluga.ru":                       {},
	"https://entry.admoblkaluga.ru/Web/Login":            {},
	"https://mimz.admoblkaluga.ru" :                      {},
	"https://mimz.admoblkaluga.ru/GzwSP/ContractsJson" :  {},
	"https://pre.admoblkaluga.ru/main/":                  {},
	"https://sadko30.admoblkaluga.ru/account/login":      {},
	"https://soc.admoblkaluga.ru":                        {},
	"https://spo.admoblkaluga.ru/security/":              {},
	"https://tender.admoblkaluga.ru":                     {},
	"https://dnmchs.ru":                                  {},
	"https://govdnr.ru":                                  {},
	"https://mer.govdnr.ru":                              {},
	"https://mid-dnr.su/ru/":                             {},
	"https://midural.ru":                                 {},
	"https://минспорт.рус":                               {},
	"https://agmo.mosreg.ru":                             {},
	"https://br.mosreg.ru/login/auth/":                   {},
	"https://dk.mosreg.ru":                               {},
	"https://easuz.mosreg.ru":                            {},
	"https://gasu.mosreg.ru":                             {},
	"https://invest.mosreg.ru":                           {},
	"https://pass.mosreg.ru/login":                       {},
	"https://login.school.mosreg.ru":                     {},
	"https://web.mail.mosreg.ru/SOGo/":                   {},
	"https://market.mosreg.ru":                           {},
	"https://rgis.mosreg.ru/v3/":                         {},
	"https://sso.mosreg.ru/signin.jsp":                   {},
	"https://support.mosreg.ru":                          {},
	"https://uslugi.mosreg.ru":                           {},
	"https://vmeste.mosreg.ru/login":                     {},
	"https://welcome.mosreg.ru":                          {},
	"http://wow.mosreg.ru/admin/":                        {},
	"https://edu-mosreg.ru":                              {},
	"https://mzdnr.ru":                                   {},
	"https://orel-region.ru":                             {},
	"https://pravdnr.ru":                                 {},
	"https://supcourt-dpr.su":                            {},
	"https://ulgov.ru":                                   {},
	"https://www.volgograd.ru":                           {},

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
	"https://wats.mirconnect.ru": {},
	"https://passport.etpgpb.ru": {},
	"https://bki-okb.ru":         {},

	// Exchanges connected to russian banks
	"https://bitokk.biz":      {},
	"https://www.netex24.net": {}, // requires cookie through redirect to fully load
	"https://flashobmen.com":  {},

	// Electronic signature services, certificate authorities, www domain names
	"https://kb.kubankredit.ru/login":     {},
	"https://structure.mil.ru":            {},
	"https://parus-s.ru":                  {},
	"https://www.icentr.ru":               {},
	"https://etp.kartoteka.ru/index.html": {},
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
	"https://fin.tektorg.ru/bg":                                   {},
	"https://fin.tektorg.ru/api/v2/reference/service_types":       {},
	"https://zakupki.tektorg.ru":                                  {},
	"https://aaa-srt-cs.lukoil.com/logon/LogonPoint/tmindex.html": {},
	"https://ar.lukoil.com/login":                                 {},
	"https://b2b.sibur.ru/pages_new_en/faces/index.jspx":          {},
	"https://onlinecontract.ru":                                   {},

	// Food delivery services
	"https://samokat.ru":    {},

	// Cinemas
	"https://karofilm.ru":   {},
	"https://www.mirage.ru": {},

	// Belgorod
	"https://bel.ru": {},
	"https://bel.ru/api/v1/platform/main_page":                 {},
	"https://belgorod.chibbis.ru/":                             {},
	"https://belgorod.igooods.ru":                              {},
	"https://belgorod.igooods.ru/api/v8/cities/identify_by_ip": {},
	"https://belgorod.farfor.ru":                               {},
	"https://belsbyt.ru":                                       {},
	"https://lk.belsbyt.ru/fiz/login":                          {},
	"https://www.go31.ru":                                      {},
	"https://www.go31.ru/api3/auth/check-login":                {},
	"https://perevozki31.ru":                                   {},

	// Others
	"http://217.12.104.100":  {},

	// Various websites by ip
	"https://213.59.254.8":   {},
	"https://213.59.197.65":  {},
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
	"https://www.belta.by":        {},
	"https://www.belarus.by/ru/":  {},
	"https://ont.by":              {}, // should be called without params
	"https://www.belnovosti.by":   {},

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
	"https://www.mid.ru":         {},

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

	// https://t.me/itarmyofukraine2022/248
	"https://digital.ac.gov.ru":       {},
	"https://platform.digital.gov.ru": {},
	"https://reestr.digital.gov.ru":   {},
	"https://sc.digital.gov.ru/home":  {},

	// https://t.me/itarmyofukraine2022/251
	"https://payanyway.ru/info/w/ru/public/welcome.htm": {},

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
	"https://wink.ru":           {},

	// https://t.me/itarmyofukraine2022/289
	"https://stream.1tv.ru/api/playlist/1tvch-v1_as_array.json": {},

	// https://t.me/itarmyofukraine2022/290
	"https://api.kontur.ru/csi/support/v1/public/integrations/zakupki": {},
	"https://auth.kontur.ru":                       {},
	"https://ca.kontur.ru":                         {},
	"https://developer.kontur.ru":                  {},
	"https://elbank.kontur.ru/AccessControl/Login": {},
	// 	"https://focus.kontur.ru":                      {}, // TODO read buffer size
	"https://focus.kontur.ru/api/lists/all": {},
	"https://pf.kontur.ru":                  {},
	"https://zakupki.kontur.ru":             {},
	"https://www.b-kontur.ru":               {},
	"https://www.kontur-extern.ru":          {},

	// https://t.me/itarmyofukraine2022/302
	"https://rov.aero": {},
}
