package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetWebsites = map[string]struct{}{
	/* Russia */

	// Propaganda
	// 	"https://tass.ru":                            {}, // stormwall
	// 	"https://tass.ru/userApi/getNewsFeed":        {}, // stormwall
	"https://vz.ru":                              {},
	"https://www.gazeta.ru":                      {},
	"https://api.riafan.ru":                      {},
	"https://mt.riafan.ru/blog/latest":           {},
	"https://pikabu.ru":                          {},
	"https://api.pikabu.ru":                      {},
	"https://www.rambler.ru":                     {},
	"https://www.rambler.ru/api/v4/personalized": {},
	// 	"https://id.rambler.ru/login-20/login": {}, // TODO fix small read buffer
	"https://news.rambler.ru":                           {},
	"https://peroxide.rambler.ru/v1/comments/clusters/": {},
	"https://rcm.rambler.ru":                            {},
	"https://zoomdecorate.rambler.ru":                   {},
	"https://mail.ru":                                   {},
	"https://ad.mail.ru/adp/":                           {},
	"https://recostream.go.mail.ru":                     {},
	"https://suggests.go.mail.ru/sg_main":               {},
	"https://news.mail.ru":                              {},
	"https://portal.mail.ru/NaviData":                   {},
	"https://rg.ru":                                     {},
	"https://360tv.ru":                                  {},
	// 	"https://47news.ru":                                 {}, // qrator
	"https://77.244.221.63":   {}, // 47news.ru
	"https://www.afanasy.biz": {},
	//  	"http://allpravda.info":                              {}, // cf
	"https://anews.com/tag/novosti/500/":                 {},
	"https://antifashist.com":                            {}, // https://antifashist.online
	"https://antimaydan.info":                            {},
	"https://argumenti.ru":                               {},
	"https://aurora.network":                             {},
	"https://balrad.ru":                                  {}, // only hu net
	"https://www.business-gazeta.ru":                     {},
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
	"https://social.dni.ru":                              {},
	"https://dnr24.com":                                  {},
	"https://dnr24.su":                                   {}, // cf
	// 	"https://dnr-life.ru":   {},
	"https://dnr-pravda.ru": {},
	// 	"https://donbasstoday.ru": {}, // stormwall
	"https://193.233.15.62":          {}, // donbasstoday.ru
	"https://dontimes.ru":            {},
	"https://dosie.su":               {}, // cf
	"https://eadaily.com/ru/dossier": {}, // cf
	"https://www.europereloaded.com": {},
	"https://evening-crimea.com":     {},
	"https://expert.ru":              {},
	"https://fine-news.ru":           {},
	// 	"https://forpostsevastopol.ru":   {},  // cf only ru net
	"https://81.200.119.206":      {},
	"https://front-novorossii.ru": {},
	// 	"http://globalwarnews.ru":             {}, // ddg
	//     "http://213.155.21.184": {}, // globalwarnews.ru
	"https://gorlovka-news.su":            {}, // cf
	"https://gtrklnr.com":                 {},
	"https://infobrics.org":               {},
	"https://infobrics.org/api/business/": {},
	// 	"https://inforeactor.ru":              {}, // airee
	"http://82.202.163.4": {}, // inforeactor.ru
	"https://infox.sg":    {},
	// 	"https://inlugansk.ru":                             {}, // cf only ru net
	"https://jpgazeta.ru":  {},
	"https://kianews24.ru": {},
	// 	"http://k-politika.ru":                             {}, // cf
	"http://91.228.154.78": {}, // k-politika.ru
	// 	"https://kuban24.tv":                               {}, // cf
	"http://82.202.160.250":                            {}, // kuban24.tv
	"https://life.ru":                                  {},
	"https://lug-info.com":                             {},
	"https://lugansk-online.su":                        {},
	"https://www.m24.ru":                               {},
	"https://cross.m24.ru/covid/frontend/web/site/rus": {},
	"https://mediarepost.ru":                           {},
	"https://www.metronews.ru":                         {},
	"https://mirtesen.ru":                              {},
	"https://www.mk-kalm.ru":                           {},
	"https://mkset.ru":                                 {},
	"http://www.moscow-post.su":                        {},
	"https://mosregtoday.ru":                           {}, // cf
	"http://178.154.213.69":                            {}, // mosregtoday.ru
	"https://www.nakanune.ru":                          {},
	// 	"https://nation-news.ru":                           {}, // airee
	"http://94.198.55.82":              {}, // nation-news.ru
	"https://newizv.ru":                {},
	"https://newizv.ru/api/v1/matters": {},
	"https://newsland.com":             {},
	"https://nvo.ng.ru":                {},
	"https://nnews.nnov.ru":            {},
	"https://newdaynews.ru":            {},
	"https://www.newc.info":            {},
	// 	"https://odnarodyna.org":                             {}, // TODO small read buffer
	"https://novorossiia.ru":   {}, // cf
	"https://87.236.16.191":    {}, // novorossiia.ru
	"https://novorossiia.info": {},
	"http://novorossy.ru":      {},
	"https://novosti.icu":      {}, // cf
	"http://novosti333.ru":     {},
	"http://webnovosti.info":   {},
	"https://nt1941.su":        {},
	"https://www.osnmedia.ru":  {},
	"https://pandoraopen.ru":   {},
	"https://plainnews.ru":     {}, // cf
	"https://pravdoryb.info":   {}, // cf
	"http://95.142.44.8":       {}, // pravdoryb.info
	// 	"https://polit.info":                  {}, // airee
	"https://152.89.217.58": {}, // polit.info
	"https://politcentr.ru": {}, // cf
	// 	"https://politexpert.net":             {}, // airee
	"http://188.127.226.136": {}, // politexpert.net
	"https://politikus.info": {}, // cf
	"https://politikus.ru":   {},
	"https://politpuzzle.ru": {},
	// 	"https://politros.com":                {}, // airee
	"http://152.89.219.221":               {}, // politros
	"http://134.0.118.81":                 {}, // ptoday.ru
	"https://rf-smi.ru":                   {},
	"https://www.ritmeurasia.org":         {},
	"https://rodina.news":                 {}, // cf
	"https://rossaprimavera.ru/feed/news": {},
	// 	"https://rusnext.ru":                  {}, // cf
	// 	"https://rusnext.ru/newslistapi.json": {}, // cf
	"http://ru-an.info":      {},
	"https://rueconomics.ru": {}, // airee
	"https://rusonline.org":  {},
	"http://ruspravda.info":  {},
	// 	"https://sevnews.info":              {},
	// 	"https://slovodel.com":              {}, // airee
	"http://94.198.50.179":              {}, // slovodel.com
	"https://smi2.ru":                   {},
	"https://polls.smi2.ru/body/1/poll": {},
	"https://sovross.ru":                {},
	"http://stringer-news.com":          {},
	"https://svpressa.ru":               {},
	"https://taurica.net":               {},
	"https://tehnowar.ru":               {},
	"https://thesaker.is":               {},
	"https://time-news.net":             {}, // cf
	"https://trmzk.ru":                  {},
	// 	"https://tvzvezda.ru":               {}, // qrator
	// 	"https://u-f.ru":                      {}, // cf
	"https://ugyalta.com":      {}, // cf
	"https://universe-tss.su":  {},
	"https://utro.ru":          {},
	"https://www.vedomosti.ru": {},
	"https://vesti-k.ru":       {},
	"https://vm.ru":            {},
	"https://voskhodinfo.su":   {},
	// 	"https://webkamerton.ru":              {}, // small read buffer
	"https://wpristav.ru":   {}, // cf
	"http://193.109.246.53": {}, // wpristav.ru
	// 	"https://x-true.info":     {}, // cf
	// 	"http://xvesti.ru":        {}, // cf
	"http://31.31.198.206": {}, // xvesti.ru
	// 	"https://ya62.ru":      {}, // ddg
	// 	"https://zanamipravda.ru": {}, // only ru net
	"https://zavtra.ru": {},

	// https://www.state.gov/russias-pillars-of-disinformation-and-propaganda-report/
	"http://www.geopolitika.ru":                           {},
	"https://katehon.com":                                 {},
	"https://novorosinform.org":                           {},
	"https://www.pravda.ru":                               {}, // cf
	"https://www.pravda.ru/ajaxed/toolbartabs/":           {}, // cf
	"https://russia-insider.com/en":                       {}, // cf
	"https://russian-faith.com":                           {},
	"http://158.69.116.70":                                {}, // russian-faith.com russia-insider.com
	"https://southfront.org":                              {},
	"https://www.strategic-culture.org":                   {},
	"https://therussophile.org":                           {},
	"https://tsargrad.tv":                                 {},
	"https://cloud.tsargrad.tv/nextcloud/index.php/login": {},
	"https://mirtesen.tsargrad.tv":                        {},
	"https://www.voltairenet.org":                         {}, // cf
	"https://65.108.13.235":                               {}, // voltairenet.org

	// Business corporations
	// 	"http://www.crimearw.ru": {},
	"https://vcs.rostec.ru": {},

	// The state
	"https://mil.ru":                  {},
	"https://vsednr.ru":               {},
	"https://dnmchs.ru":               {},
	"http://www.fsb.ru":               {},
	"https://govdnr.ru":               {},
	"https://mer.govdnr.ru/index.php": {},
	"https://mid-dnr.su/ru/":          {},
	"https://mzdnr.ru":                {},
	"https://pravdnr.ru":              {},

	// Embassies
	// Do not duplicate the IPs: https://github.com/erkexzcx/stoppropaganda/pull/110#issuecomment-1059960305
	"https://montreal.mid.ru": {},
	"https://belarus.mid.ru":  {},

	// Belgorod
	"https://bel.ru": {},
	"https://bel.ru/api/v1/platform/main_page":  {}, // cf
	"https://5.188.73.213":                      {}, // www.bel.ru
	"https://belsbyt.ru":                        {},
	"https://lk.belsbyt.ru/fiz/login":           {},
	"https://www.go31.ru":                       {},
	"https://www.go31.ru/api3/auth/check-login": {},

	/* BELARUS */

	// by gov
	"https://president.gov.by/ru":   {},
	"https://mfa.gov.by":            {},
	"https://russia.mfa.gov.by/ru/": {},

	// by media
	"https://www.belta.by":       {},
	"https://www.belarus.by/ru/": {},
	"https://ont.by":             {}, // should be called without params
	"https://www.belnovosti.by":  {},

	// Still operating in Russia
	// https://www.yalerussianbusinessretreat.com/
	"https://burgerkingrus.ru":                                      {},
	"https://burgerkingrus.ru/api-web-front/api/v3/restaurant/list": {},
	"https://mic.burgerking.ru/mifs/user/login.jsp":                 {},
	"https://sd.burgerking.ru/HEAT/SaaSExternalSessionRenew.aspx":   {},
	"https://delonghi.ru":                                           {},
	// 	"https://ecco.ru":                                               {}, // only ru net
	"https://95.215.182.186": {}, // ecco.ru
	"https://www.ehrmann.ru": {},
	// 	"http://flowserve-industry.ru":                                  {}, // ddg
	// 	"https://www.laredoute.ru":                                {}, // only ru net
	// 	"https://www.laredoute.ru/producthelper/getproducts.aspx": {},
	"https://mega.ru/offers/":              {}, // Ingka
	"https://modshairrussia.ru":            {},
	"https://mymodshair.ru":                {},
	"https://myzte.ru":                     {},
	"https://natr.ru":                      {},
	"https://papajohns.ru/moscow":          {},
	"https://api.papajohns.ru/slider/list": {},
	"https://subway.ru":                    {},

	// https://t.me/itarmyofukraine2022/230
	"https://www.mid.ru": {},

	// https://t.me/itarmyofukraine2022/215
	"https://www.chechnya.online": {},

	// https://t.me/itarmyofukraine2022/289
	"https://stream.1tv.ru/api/playlist/1tvch-v1_as_array.json": {},
}
