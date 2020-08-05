<template>
    <div id="kakao-login-btn">
    </div>
</template>

<script>
    export default {
        name: "AuthKakao",
        methods: {
            kakaoLogin() {
                Kakao.Auth.authorize({
                    redirectUri: `${window.location.origin}/auth/kakao`
                })
            }
        },
        mounted() {
            const script = document.createElement('script')
            script.src = 'https://developers.kakao.com/sdk/js/kakao.min.js'
            script.onload = (() => {
                Kakao.init('eaae225fb5f9be2ef750cc65cf9822a7')
                Kakao.Auth.createLoginButton({
                    container: '#kakao-login-btn',
                    success: function(authObj) {
                        Kakao.API.request({
                            url: '/v2/user/me',
                            success: function (res) {
                                //수정 필요
                                this.$store.dispatch('loginWithKakao', res)
                            },
                            fail: function (error) {
                                alert('사용자 정보를 가져오는 데에 실패하였습니다. 다시 시도해 주세요.');
                                window.location.replace('/login');
                            },
                        })
                    }
                })
            })
            document.body.appendChild(script)
        }
    }
</script>

<style scoped>

</style>