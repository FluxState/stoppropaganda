package targets

// Source: https://twitter.com/FedorovMykhailo/status/1497642156076511233

var TargetWebsites = map[string]struct{}{
	/* Russia */

	// Propaganda
	"https://12-kanal.ru":   {}, //
	"https://360tv.ru":      {}, // q | should be called without query
	"https://47news.ru":     {}, // qrator
	"https://77.244.221.63": {}, // 47news.ru
	"https://78.ru":         {},
	// 	"https://www.afanasy.biz":                           {}, // cf only ru net
	//  	"http://allpravda.info":                              {}, // cf
	"https://anews.com": {},
	// 	"https://antifashist.com": {}, // https://antifashist.online | only ru net
	"https://antimaydan.info": {}, // should be called without query
	"https://argumenti.ru":    {},
	"https://aurora.network":  {}, // should be called without query
	// 	"https://balrad.ru":                  {}, // cf only hu net
	"https://bel.ru": {}, // cf
	"https://bel.ru/api/v1/platform/main_page": {}, // cf
// 	"https://bk55.ru":                          {}, // sw
	"https://bloknot-rostov.ru":                {},
	"https://www.bnkomi.ru":                    {},
	// 	"https://www.business-gazeta.ru":                     {}, // qrator
	"https://buzina.org":        {},
	"https://www.c-inform.info": {},
	// 	"https://cherkessk09.ru":                             {}, // dg
	"https://crimea-news.com":  {},
	"https://crimeapress.info": {},
	"https://comitet.su":       {},
	"https://compromat.group":  {},
	"https://cont.ws":          {},
	"https://dan-news.info":    {},
	"https://dni.ru":           {},
	"https://social.dni.ru":    {},
	"https://dnr24.com":        {},
	"https://dnr24.su":         {}, // cf
	// 	"https://dnr-life.ru":   {},
	"https://dnr-pravda.ru":   {}, // sw
	"https://donbasstoday.ru": {}, // stormwall | should be called without query
	"https://www.donnews.ru":  {},
	"https://dontimes.ru":     {},
	// 	"https://dosie.su":               {}, // cf
	"https://eadaily.com/ru/dossier": {}, // cf
	"https://www.europereloaded.com": {},
	"https://evening-crimea.com":     {}, // should be called without query
	"https://expert.ru":              {},
	"https://fine-news.ru":           {},
	// 	"https://forpostsevastopol.ru":   {}, // cf only ru net
	// 	"https://front-novorossii.ru": {},
	"https://www.gazeta.ru": {},
	// 	"http://globalwarnews.ru":             {}, // dg
	"https://gorlovka-news.su": {}, // cf
	"https://gtrklnr.com":      {},
	// 	"https://inforeactor.ru":              {}, // airee
	// 	"https://inlugansk.ru":                             {}, // cf only ru net
	// 	"https://jpgazeta.ru": {}, // airee
	// 	"https://kianews24.ru": {}, // dg
	// 	"http://k-politika.ru": {}, // cf
	// 	"https://kuban24.tv":        {}, // cf should be called without query
	"https://life.ru":                   {},
	"https://lug-info.com":              {},
	"https://lugansk-online.su":         {},
	"https://www.m24.ru":                {},
	"https://mail.ru":                   {},
	"https://news.mail.ru":              {},
	"https://mediarepost.ru":            {},
	"https://www.metronews.ru":          {},
	"https://mirbelogorya.ru/index.php": {},
	"https://mirtesen.ru":               {},
	// 	"https://mkset.ru":                                 {}, // cf only ru net
	// 	"https://www.mordovmedia.ru": {}, // cf
	"http://www.moscow-post.su": {}, // sw
	"https://mosregtoday.ru":    {}, // cf
	"https://www.nakanune.ru":   {},
	// 	"https://nation-news.ru":                           {}, // airee
	"https://newizv.ru":                {},
	"https://newizv.ru/api/v1/matters": {},
	"https://newsland.com":             {},
	"https://nvo.ng.ru":                {},
	"https://newdaynews.ru":            {},
	// 	"https://www.newc.info":            {},
	"https://novorossiia.ru":   {}, // cf
	"https://novorossiia.info": {}, // should be called without query
	"http://novorossy.ru":      {},
	"https://novosti.icu":      {}, // cf
	// 	"http://novosti333.ru":     {},
	// 	"http://webnovosti.info":   {},
	"https://nt1941.su": {},
	// 	"https://odnarodyna.org":                             {}, // TODO small read buffer
	// 	"https://www.osnmedia.ru":  {}, // qrator
	"https://pandoraopen.ru": {}, // should be called without query
	// 	"https://pg11.ru":                       {}, // only ru net
	"https://pikabu.ru":                     {},
	"https://pizzahut.ru":                   {},
	"https://pizzahut.ru/ajax/menu-groups/": {},
	"https://plainnews.ru":                  {}, // cf
	// 	"https://polit.info":                  {}, // airee
	// 	"https://politcentr.ru": {}, // cf
	// 	"https://politexpert.net":             {}, // airee
	"https://politikus.info": {}, // cf
	"https://politikus.ru":   {},
	"https://politpuzzle.ru": {},
	// 	"https://politros.com":                {}, // airee
	// 	"https://pravdoryb.info":              {}, // cf
	"https://www.rambler.ru":                      {},
	"https://www.rambler.ru/api/v4/personalized/": {},
	// 	"https://id.rambler.ru/login-20/login": {}, // TODO fix small read buffer
	"https://news.rambler.ru":                              {},
	"https://peroxide.rambler.ru/v1/tops/common/clusters/": {},
	"https://rcm.rambler.ru":                               {},
	"https://rf-smi.ru":                                    {},
	"https://rg.ru":                                        {},
	"https://riafan.ru":                                    {}, // q
	"https://riamo.ru":                                     {},
	"https://www.ritmeurasia.org":                          {},
	"https://rodina.news":                                  {}, // cf
	"https://rossaprimavera.ru/feed/news":                  {},
	"https://rostovgazeta.ru":                              {}, // cf
	// 	"https://rusnext.ru":                  {}, // cf
	// 	"https://rusnext.ru/newslistapi.json": {}, // cf
	// 	"https://ru-an.info":     {}, // cf
	// 	"https://rueconomics.ru": {}, // airee
	"https://rusonline.org": {},
	// 	"http://ruspravda.info":  {}, // cf
	"https://russia58.tv/projects/zamir/": {},
	// 	"https://slovodel.com":              {}, // airee
	"https://smi2.ru":          {},
	"https://sovross.ru":       {},
	"http://stringer-news.com": {}, // redirect with set-cookie
	"https://svpressa.ru":      {},
	// 	"https://tass.ru":                            {}, // sw
	// 	"https://tass.ru/userApi/getNewsFeed":        {}, // sw
	"https://taurica.net":      {},
	"https://tehnowar.ru":      {},
	"https://thesaker.is":      {},
	"https://time-news.net":    {}, // cf
	"https://trmzk.ru":         {},
	"https://tvr24.tv/top-100": {},
	// 	"https://tvzvezda.ru":               {}, // qrator
	// 	"https://u-f.ru":                      {}, // cf
	// 	"https://ufanotes.ru":      {},
	"https://ugyalta.com": {}, // cf
	// 	"https://universe-tss.su":  {},
	"https://utro.ru":          {},
	"https://www.vedomosti.ru": {},
	// 	"https://vesti-k.ru":       {}, // only ru net
	"https://vestnickprimanich.ru": {},
	"https://vkirove.ru":           {},
	// 	"https://vm.ru":            {}, // dg | should be called without query
	"https://voskhodinfo.su": {},
	// 	"https://webkamerton.ru":              {}, // small read buffer
	"https://wpristav.com": {}, // (wpristav.ru cf)
	// 	"https://x-true.info":     {}, // cf
	// 	"http://xvesti.ru":        {}, // cf
	// 	"https://ya62.ru":      {}, // dg
	// 	"https://zanamipravda.ru": {}, // only ru net
	// 	"https://zavtra.ru": {}, // cf | only ru net

	// https://www.state.gov/russias-pillars-of-disinformation-and-propaganda-report/
	"https://www.geopolitika.ru": {},
	"https://katehon.com":        {},
	"https://novorosinform.org":  {},
	// 	"https://www.pravda.ru":                     {}, // cf
	// 	"https://www.pravda.ru/ajaxed/toolbartabs/": {}, // cf
	"https://russia-insider.com/en":     {}, // cf
	"https://russian-faith.com":         {},
	"https://southfront.org":            {},
	"https://www.strategic-culture.org": {}, // redirect with cookie
	"https://therussophile.org":         {},
	"https://tsargrad.tv":               {},
	"https://mirtesen.tsargrad.tv":      {},
	// 	"https://www.voltairenet.org":       {}, // fi hetzner | cf | should be called without params
	"https://vz.ru": {},

	// Business corporations
	// 	"http://www.crimearw.ru": {},

	// The state
	"https://mil.ru":    {}, // should be called without query
	"https://vsednr.ru": {},
	"https://dnmchs.ru": {},
	// 	"http://www.fsb.ru":               {}, // only ru net
	"https://govdnr.ru":               {}, // only ru net?
	"https://mer.govdnr.ru/index.php": {},
	"https://mid-dnr.su/ru/":          {},
	// 	"https://mzdnr.ru":                {}, // cf | only ru net
	"https://pravdnr.ru": {}, // only ru net?

	// Embassies
	// Do not duplicate the IPs: https://github.com/erkexzcx/stoppropaganda/pull/110#issuecomment-1059960305
	"https://montreal.mid.ru": {},
	"https://belarus.mid.ru":  {},

	/* BELARUS */

	// by gov
	"https://president.gov.by/ru":   {}, // should be called without query
	"https://mfa.gov.by":            {},
	"https://russia.mfa.gov.by/ru/": {},

	// by media
	"https://www.belta.by":       {},
	"https://www.belarus.by/ru/": {},
	"https://ont.by":             {}, // should be called without params
	"https://www.belnovosti.by":  {},

	// Still operating in Russia
	// https://www.yalerussianbusinessretreat.com/
	"https://ames.ru":          {}, // Nor-Maali
	"https://burgerkingrus.ru": {},
	"https://burgerkingrus.ru/api-web-front/api/v3/restaurant/list": {},
	"https://mic.burgerking.ru/mifs/user/login.jsp":                 {},
	"https://sd.burgerking.ru/HEAT/SaaSExternalSessionRenew.aspx":   {},
	"https://www.clarins.ru":                                        {}, // cf
	"http://91.189.114.29":                                          {}, // clarins.ru
	// 	"https://www.colins.ru":                                         {}, // only ru net
	// 	"https://www.xn----7sbbaeo6behzdm8c.xn--p1ai":                   {}, // doka | qrator
	"https://delonghi.ru": {}, // should be called without query
	// 	"https://ecco.ru":                                               {}, // only ru net
	"https://obyv-online.ru": {}, // ecco
	"https://www.ehrmann.ru": {},
	// 	"https://www.ela-container.ru":                                  {}, // one global server
	// 	"https://www.etam.ru":                                           {}, // cf server in fr
	"https://fenzirussia.com": {},
	// 	"https://fischer-russia.com": {}, // only ru net
	// 	"http://flowserve-industry.ru":                                  {}, // dg
	"https://foraco.ru":            {},
	"https://www.foreverliving.ru": {},
	"https://frigoglass-shop.ru":   {},
	// 	"http://hardrockcafe.ru":                                 {},
	// 	"https://hardrockcafespb.ru":                             {}, // cf
	"https://hoffmann-group.ru":                              {},
	"https://www.kfc.ru":                                     {},
	"https://api.kfc.com/api/store/v2/store.get_restaurants": {},
	// 	"https://www.laredoute.ru":                                {}, // cd
	// 	"https://www.laredoute.ru/producthelper/getproducts.aspx": {},
	"https://legrand.ru": {},
	// 	"https://www.leptos-estates.ru":                           {}, // server in uk
	"https://www.makita.ru":                {},
	"https://mega.ru/offers/":              {}, // Ingka
	"https://modshairrussia.ru":            {},
	"https://moutairussia.com":             {},
	"https://mymodshair.ru":                {},
	"https://myzte.ru":                     {},
	"https://natr.ru":                      {},
	"https://www.obocom.ru":                {},
	"https://oppo.ru":                      {},
	"https://papajohns.ru/moscow":          {},
	"https://api.papajohns.ru/slider/list": {},
	"https://projahn-rus.ru":               {},
	"https://projahn-russia.ru":            {},
	"https://projahn.su":                   {},
	"https://ipoteka.raiffeisen.ru/signin": {},
	"https://auth.ipoteka.raiffeisen.ru/realms/origin/protocol/openid-connect/auth": {},
	"https://partner.ipoteka.raiffeisen.ru/kclogin":                                 {},
	"https://online.raiffeisen.ru/login/main":                                       {},
	"https://skidki.raiffeisen.ru":                                                  {},
	"https://sbarro-pizza.ru":                                                       {},
	"https://sbdr.ru":                                                               {},
	// 	"https://www.sgs.ru":                                                            {}, // server in nl
	// 	"https://www.storck.ru/ru/":                                                     {}, // server in de
	"https://subway.ru":      {},
	"https://swiss-krono.ru": {},
	// 	"https://triumphlingerie.ru":  {},
	// 	"https://www.wienerberger.ru": {}, // server on aws
	// 	"https://www.zippo.ru": {}, // only ru net

	// https://t.me/itarmyofukraine2022/230
	"https://www.mid.ru": {},

	// https://t.me/itarmyofukraine2022/215
	"https://www.chechnya.online": {},
}
