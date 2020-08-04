<template>
    <div class="kakaoCallBack">
        <h3> 잠시만 기다려주세요 </h3>
    </div>
</template>

<script>
    export default {
        name: "KakaoCallBack",
        data() {
            return {
            }
        },
        created() {
            let params = getUrlParams()

            if (params.error) {
                alert('로그인이 취소되었거나 만 14세 미만 보호자의 동의가 실패하였습니다. 다시 시도해 주세요.')
                window.location.replace('/login')
                return
            }
            if (!params.code) {
                alert('인증 코드를 가져오는 데에 에러가 발생했습니다: \n다시 시도해 주세요.')
                window.location.replace('/login')
                return
            }
            this.$axios.post(`https://kauth.kakao.com/oauth/token?grant_type=authorization_code&client_id=f870f92ceeb2c80da32b754d287c046a&redirect_uri=http://localhost:8080/auth/kakao&client_secret=YkCo7LlVRHrXPON9rxmoCkvtwWXKV0Xj&code=`+params.code)
                .then((res) => {
                    this.$axios.get(`https://kauth.kakao.com/v2/user/me`, {
                        headers: {
                            'Authorization': 'Bearer ' + res.data.access_token,
                            'Content-type': 'application/x-www-form-urlencoded;charset=utf-8'
                        }
                    }).then((res) => {
                        console.log(res)
                        alert(res.data.id + ' ' + res.data.kakao_account.profile.nickname)
                        //this.$store.dispatch('loginWithKakao', ).then(window.location.replace('/'))
                    }).catch((err) => {
                        alert(err)
                        window.location.replace('/login')
                    })
                })
                .catch((err) => {
                    alert(err)
                    window.location.replace('/login')
                })
        },
    }
    function getUrlParams() {
        var params = {};
        window.location.search.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(str, key, value) { params[key] = value; });
        return params;
    }
</script>

<style scoped>

</style>