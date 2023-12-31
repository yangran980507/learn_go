// Package validators 存放自定义规则及验证器
package validators

import (
	"gohub/pkg/captcha"
	"gohub/pkg/verifycode"
)

// ValidateCaptcha 自定义规则，验证『图片验证码』
func ValidateCaptcha(captchaID, captchaAnswer string,
	errs map[string][]string) map[string][]string {
	if ok := captcha.NewCaptcha().VerifyCaptcha(captchaID, captchaAnswer); !ok {
		//delete(errs, "captcha_answer")
		errs["captcha_answer"] = append(errs["captcha_answer"], "图片验证码错误")
	}
	return errs
}

// ValidatePasswordConfirm 检查两次密码事否正确
func ValidatePasswordConfirm(password, passwordConfirm string,
	errs map[string][]string) map[string][]string {
	if password != passwordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "俩次输入密码不正确")
	}
	return errs
}

// ValidateVerifyCode 自定义规则，验证『手机/邮箱验证码』
func ValidateVerifyCode(key, answer string,
	errs map[string][]string) map[string][]string {
	if ok := verifycode.NewVerifyCode().CheckAnswer(key, answer); !ok {
		errs["verify_code"] = append(errs["verify_code"], "验证码错误")
	}
	return errs
}
