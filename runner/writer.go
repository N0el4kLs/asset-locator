package runner

import (
	"bytes"
	"strings"

	"github.com/N0el4kLs/asset-locator/pkg/sources/weight/providers"

	"github.com/logrusorgru/aurora"
	"github.com/projectdiscovery/gologger"
)

var (
	NO_ICP_FOUND = "no icp found"
)

type Result struct {
	// Target is input item to do weight scan or icp scan
	// It might be in the format: IP, Domain or URL
	Target string
	// ICP is the ICP inform of the target
	ICP string
	// Weight is the weight of the target
	Weight providers.WeightLevel
	// tType is the type of the input item
	// It might be IP, Domain or URL
	tType TargetType
	// vValue extract ip or domain from the target
	vValue string
}

func NewResult(target string) *Result {
	tValue, tType := parseTarget(target)

	return &Result{
		Target: target,
		Weight: providers.ErrorLevel,
		ICP:    NO_ICP_FOUND,
		tType:  tType,
		vValue: tValue,
	}
}

type Writer struct {
	StdWriter StdWriter
}

type StdWriter struct {
	options         Options
	withNoColor     bool          // color option
	showOnly        bool          // only show result option
	ColorClient     aurora.Aurora // color client
	validWeightOnly bool          // only show valid weight
}

var (
	ExclusiveSRC = []string{
		"上海合合信息科技股份有限公司",
		"北京粉笔蓝天科技有限公司",
		"乐信集团",
		"shein",
		"广汽本田汽车有限公司",
		"申通快递有限公司",
		"奇安信集团",
		"信也科技集团",
		"人民教育出版社",
		"中国外汇交易中心",
		"华住酒店",
		"上海亘岩网络科技有限公司",
		"北京北森云计算股份有限公司",
		"瓜子二手车",
		"千寻位置网络有限公司",
		"上海泛微网络科技股份有限公司",
		"通达OA",
		"58到家",
		"水滴公司",
		"广东堡塔安全技术有限公司",
		"山东省残疾人联合会",
		"宝宝树",
		"搜狐",
		"翼支付",
		"同程旅行",
		"汽车之家",
		"北京易车信息技术有限公司",
		"环球时报在线（北京）文化传播有限公司",
		"广州快猫科技有限公司",
		"浙江物产信息技术有限公司",
		"诺安基金管理有限公司",
		"厦门航空",
		"北京值得买科技股份有限公司",
		"51信用卡",
		"安徽甜心互娱网络科技有限公司",
		"好大夫在线",
		"泓德基金管理有限公司",
		"江苏方天电力技术有限公司",
		"安徽省刀锋网络科技有限公司",
		"杭州古北电子科技有限公司",
		"上海贝锐信息科技股份有限公司",
		"同盾科技有限公司",
		"嘀嗒出行",
		"广东速狐信息科技有限公司",
		"北京神州汽车租赁有限公司",
		"上海付费通信息服务有限公司",
		"杭州恒业网络信息有限公司",
		"帆软软件有限公司",
		"丁香园",
		"北京网元圣唐娱乐科技有限公司",
		"华夏基金",
		"货讯通科技",
		"中国电信股份有限公司江西分公司",
		"中国电信股份有限公司上海分公司",
		"国海证券",
		"石家庄哆吧网络科技有限公司",
	}
)

func DefaultCallout(rst Result) {
	buf := bytes.Buffer{}
	buf.WriteString(rst.Target)

	buf.WriteString(" [")
	buf.WriteString(rst.Weight.ToString())
	buf.WriteString("]")

	buf.WriteString(" [")
	buf.WriteString(rst.ICP)
	buf.WriteString("]")

	if v, ok := inExclusiveSrc(rst.Target); ok {
		buf.WriteString(" [")
		buf.WriteString(v)
		buf.WriteString("]")
	}

	gologger.Silent().Msgf("%s\n", buf.String())
}

func inExclusiveSrc(t string) (string, bool) {
	for _, src := range ExclusiveSRC {
		if strings.Contains(t, src) {
			return src, true
		}
	}
	return "", false
}
