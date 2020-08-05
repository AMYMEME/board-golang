<template>
    <div id="kakao-login-btn">
    </div>
</template>

<script>
    export default {
        name: "AuthKakao",
        methods: {
            kakaoLogin(store) {
                Kakao.init('eaae225fb5f9be2ef750cc65cf9822a7')
                Kakao.Auth.createLoginButton({
                    container: '#kakao-login-btn',
                    success: function() {
                        Kakao.API.request({
                            url: '/v2/user/me',
                            success: function (res) {
                                store.dispatch('loginWithKakao', res)
                                    .then(() =>  this.backRedirect())
                            },
                            fail: function () {
                                alert('사용자 정보를 가져오는 데에 실패하였습니다. 다시 시도해 주세요.');
                                window.location.replace('/login');
                            },
                        })
                    }
                })
            },
            backRedirect() {
                let redirectURI;
                redirectURI = decodeURIComponent(getUrlParams().redirect);
                if (redirectURI === 'undefined') {
                    this.$router.push('/');
                } else {
                    this.$router.push(redirectURI)
                }
            },
        },
        mounted() {
            const script = document.createElement('script')
            script.src = 'https://developers.kakao.com/sdk/js/kakao.min.js'
            script.onload = (() => {
                this.kakaoLogin(this.$store)
            })
            document.body.appendChild(script)
        }
    }
    function getUrlParams() {
        var params = {};
        window.location.search.replace(/[?&]+([^=&]+)=([^&]*)/gi, function(str, key, value) { params[key] = value; });
        return params;
    }
</script>

<style scoped>

</style>