package oygo

// BankCode value
type BankCode string

const (
	BankBRI                        BankCode = "002"
	BankMandiri                    BankCode = "008"
	BankBNI                        BankCode = "009"
	BankDanamon                    BankCode = "011"
	BankPermata                    BankCode = "013"
	BankBCA                        BankCode = "014"
	BIIMaybank                     BankCode = "016"
	BankPanin                      BankCode = "019"
	CIMBNiaga                      BankCode = "022"
	BankUOBINDONESIA               BankCode = "023"
	BankOCBCNISP                   BankCode = "028"
	CITIBANK                       BankCode = "031"
	BankWinduKentjanaInternational BankCode = "036"
	BankARTHAGRAHA                 BankCode = "037"
	BankTOKYOMITSUBISHIUFJ         BankCode = "042"
	BankDBS                        BankCode = "046"
	Standard                       BankCode = "050"
	BankCAPITAL                    BankCode = "054"
	ANZIndonesia                   BankCode = "061"
	BankOFCHINA                    BankCode = "069"
	BankBumiArta                   BankCode = "076"
	BankHSBC                       BankCode = "087"
	BankAntardaerah                BankCode = "088"
	BankRabobank                   BankCode = "089"
	BankJTRUSTINDONESIA            BankCode = "095"
	BankMAYAPADA                   BankCode = "097"
	BankJawaBarat                  BankCode = "110"
	BankDKI                        BankCode = "111"
	BankBPDDIY                     BankCode = "112"
	BankJATENG                     BankCode = "113"
	BankJatim                      BankCode = "114"
	BankJambi                      BankCode = "115"
	BankAceh                       BankCode = "116"
	BankSUMUT                      BankCode = "117"
	BankNAGARI                     BankCode = "118"
	BankRiau                       BankCode = "119"
	BankSUMSELBABEL                BankCode = "120"
	BankLampung                    BankCode = "121"
	BankKALSEL                     BankCode = "122"
	BankKALBAR                     BankCode = "123"
	BankBPDKaltim                  BankCode = "124"
	BankBPDKalteng                 BankCode = "125"
	BankSULSELBAR                  BankCode = "126"
	BankSulut                      BankCode = "127"
	BankNTB                        BankCode = "128"
	BankBPDBali                    BankCode = "129"
	BankNTT                        BankCode = "130"
	BankMaluku                     BankCode = "131"
	BankBPDPapua                   BankCode = "132"
	BankSULTENG                    BankCode = "134"
	BankSultra                     BankCode = "135"
	BankBANTEN                     BankCode = "137"
	BankNusantaraParahyangan       BankCode = "145"
	BankOfIndiaIndonesia           BankCode = "146"
	BankMuamalat                   BankCode = "147"
	BankMestika                    BankCode = "151"
	BankSHINHAN                    BankCode = "152"
	BankSinarmas                   BankCode = "153"
	BankMaspion                    BankCode = "157"
	BankGanesha                    BankCode = "161"
	BankICBC                       BankCode = "164"
	BankQNBIndonesia               BankCode = "167"
	BankBTN                        BankCode = "200"
	BankWooriSaudara               BankCode = "212"
	BankBTPN                       BankCode = "213"
	BankVictoriaSyariah            BankCode = "405"
	BankJabarBantenSyariah         BankCode = "425"
	BankMega                       BankCode = "426"
	BankBukopin                    BankCode = "441"
	BankSyariahIndonesia           BankCode = "451"
	BankJasaJakarta                BankCode = "472"
	BankKEBHANA                    BankCode = "484"
	BankMNCINTERNATIONAL           BankCode = "485"
	BankNeoCommerce                BankCode = "490"
	BankRakyatIndonesiaAGRONIAGA   BankCode = "494"
	BankSBIIndonesia               BankCode = "498"
	BankRoyal                      BankCode = "501"
	BankNationalNobu               BankCode = "503"
	BankMEGASYARIAH                BankCode = "506"
	BankINA                        BankCode = "513"
	BankPANINSYARIAH               BankCode = "517"
	PRIMAMASTERBANK                BankCode = "520"
	BankSYARIAHBUKOPIN             BankCode = "521"
	BankSahabatSampoerna           BankCode = "523"
	BankDINAR                      BankCode = "526"
	BankKESEJAHTERAANEKONOMI       BankCode = "535"
	BankBCASYARIAH                 BankCode = "536"
	BankJago                       BankCode = "542"
	BankBTPNSYARIAH                BankCode = "547"
	BankMULTIARTASENTOSA           BankCode = "548"
	BankMayora                     BankCode = "553"
	BankINDEX                      BankCode = "555"
	CNB                            BankCode = "559"
	BankMANTAP                     BankCode = "564"
	BankVICTORIAINTL               BankCode = "566"
	HARDA                          BankCode = "567"
	BPRKS                          BankCode = "688"
	IBK                            BankCode = "945"
	BankCTBCIndonesia              BankCode = "949"
	BankCOMMONWEALTH               BankCode = "950"
	OVO                            BankCode = "ovo"
	LinkAja                        BankCode = "linkaja"
	Dana                           BankCode = "dana"
	Gopay                          BankCode = "gopay"
	ShopeePayEwallet               BankCode = "shopeepay_ewallet"
	LinkAjaEwallet                 BankCode = "linkaja_ewallet"
)

// AllBankCode : Get All available PaymentType
var AllBankCode = map[BankCode]string{
	BankBRI:                        `Bank BRI`,
	BankMandiri:                    `Bank Mandiri`,
	BankBNI:                        `Bank BNI`,
	BankDanamon:                    `Bank Danamon`,
	BankPermata:                    `Bank Permata`,
	BankBCA:                        `Bank BCA`,
	BIIMaybank:                     `BII Maybank`,
	BankPanin:                      `BankP anin`,
	CIMBNiaga:                      `CIMB Niaga`,
	BankUOBINDONESIA:               `Bank UOB INDONESIA`,
	BankOCBCNISP:                   `Bank OCBCNISP`,
	CITIBANK:                       `CITI BANK`,
	BankWinduKentjanaInternational: `Bank Windu Kentjana International`,
	BankARTHAGRAHA:                 `Bank ARTHAGRAHA`,
	BankTOKYOMITSUBISHIUFJ:         `Bank TOKYOMITSUBISHIUFJ`,
	BankDBS:                        `Bank DBS`,
	Standard:                       `Standard`,
	BankCAPITAL:                    `Bank CAPITAL`,
	ANZIndonesia:                   `ANZ Indonesia`,
	BankOFCHINA:                    `Bank OFCHINA`,
	BankBumiArta:                   `Bank BumiArta`,
	BankHSBC:                       `Bank HSBC`,
	BankAntardaerah:                `Bank Antardaerah`,
	BankRabobank:                   `Bank Rabobank`,
	BankJTRUSTINDONESIA:            `Bank JTRUST INDONESIA`,
	BankMAYAPADA:                   `Bank MAYAPADA`,
	BankJawaBarat:                  `Bank JawaBarat`,
	BankDKI:                        `Bank DKI`,
	BankBPDDIY:                     `Bank BPDDIY`,
	BankJATENG:                     `Bank JATENG`,
	BankJatim:                      `Bank Jatim`,
	BankJambi:                      `Bank Jambi`,
	BankAceh:                       `Bank Aceh`,
	BankSUMUT:                      `Bank SUMUT`,
	BankNAGARI:                     `Bank NAGARI`,
	BankRiau:                       `Bank Riau`,
	BankSUMSELBABEL:                `Bank SUMSELBABEL`,
	BankLampung:                    `Bank Lampung`,
	BankKALSEL:                     `Bank KALSEL`,
	BankKALBAR:                     `Bank KALBAR`,
	BankBPDKaltim:                  `Bank BPD Kaltim`,
	BankBPDKalteng:                 `Bank BPD Kalteng`,
	BankSULSELBAR:                  `Bank SULSELBAR`,
	BankSulut:                      `Bank Sulut`,
	BankNTB:                        `Bank NTB`,
	BankBPDBali:                    `Bank BPDBali`,
	BankNTT:                        `Bank NTT`,
	BankMaluku:                     `Bank Maluku`,
	BankBPDPapua:                   `Bank BPDPapua`,
	BankSULTENG:                    `Bank SULTENG`,
	BankSultra:                     `Bank Sultra`,
	BankBANTEN:                     `Bank BANTEN`,
	BankNusantaraParahyangan:       `Bank Nusantara Parahyangan`,
	BankOfIndiaIndonesia:           `Bank Of India Indonesia`,
	BankMuamalat:                   `Bank Muamalat`,
	BankMestika:                    `Bank Mestika`,
	BankSHINHAN:                    `Bank SHINHAN`,
	BankSinarmas:                   `Bank Sinarmas`,
	BankMaspion:                    `Bank Maspion`,
	BankGanesha:                    `Bank Ganesha`,
	BankICBC:                       `Bank ICBC`,
	BankQNBIndonesia:               `Bank QNB Indonesia`,
	BankBTN:                        `Bank BTN`,
	BankWooriSaudara:               `Bank Woori Saudara`,
	BankBTPN:                       `Bank BTPN`,
	BankVictoriaSyariah:            `Bank Victoria Syariah`,
	BankJabarBantenSyariah:         `Bank Jabar Banten Syariah`,
	BankMega:                       `Bank Mega`,
	BankBukopin:                    `Bank Bukopin`,
	BankSyariahIndonesia:           `Bank Syariah Indonesia`,
	BankJasaJakarta:                `Bank JasaJakarta`,
	BankKEBHANA:                    `Bank KEBHANA`,
	BankMNCINTERNATIONAL:           `Bank MNCINTERNATIONAL`,
	BankNeoCommerce:                `Bank NeoCommerce`,
	BankRakyatIndonesiaAGRONIAGA:   `Bank RakyatIndonesiaAGRONIAGA`,
	BankSBIIndonesia:               `Bank SBIIndonesia`,
	BankRoyal:                      `Bank Royal`,
	BankNationalNobu:               `Bank NationalNobu`,
	BankMEGASYARIAH:                `Bank MEGASYARIAH`,
	BankINA:                        `Bank INA`,
	BankPANINSYARIAH:               `Bank PANINSYARIAH`,
	PRIMAMASTERBANK:                `PRIM AMASTERBANK`,
	BankSYARIAHBUKOPIN:             `Bank SYARIAHBUKOPIN`,
	BankSahabatSampoerna:           `Bank SahabatSampoerna`,
	BankDINAR:                      `Bank DINAR`,
	BankKESEJAHTERAANEKONOMI:       `Bank KESEJAHTERAANEKONOMI`,
	BankBCASYARIAH:                 `Bank BCASYARIAH`,
	BankJago:                       `Bank Jago`,
	BankBTPNSYARIAH:                `Bank BTPNSYARIAH`,
	BankMULTIARTASENTOSA:           `Bank MULTIARTASENTOSA`,
	BankMayora:                     `Bank Mayora`,
	BankINDEX:                      `Bank INDEX`,
	CNB:                            `CNB`,
	BankMANTAP:                     `Bank MANTAP`,
	BankVICTORIAINTL:               `Bank VICTORIAINTL`,
	HARDA:                          `HARDA`,
	BPRKS:                          `BPRKS`,
	IBK:                            `IBK`,
	BankCTBCIndonesia:              `Bank CTBC Indonesia`,
	BankCOMMONWEALTH:               `Bank COMMONWEALTH`,
	OVO:                            `OVO`,
	LinkAja:                        `LinkAja`,
	Dana:                           `Dana`,
	Gopay:                          `Gopay`,
}
