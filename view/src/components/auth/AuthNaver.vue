<template>
    <div id="naverIdLogin" class="mt-4"></div>
</template>

<script>
    export default {
        name: "AuthNaver",
        methods: {
            naverLogin(store) {
                const naverLogin = new naver.LoginWithNaverId(
                    {
                        clientId: `_HXiSkrZSfB9vq94TsZC`,
                        callbackUrl: `http://localhost:8080/login`,
                        isPopup: false,
                        loginButton: {color: "green", type: 3, height: 48},
                        callbackHandle: true
                    }
                )
                naverLogin.init()

                var naverLoginBtn = document.getElementById('naverIdLogin');
                naverLoginBtn.addEventListener('click', function () {
                    naverLogin.getLoginStatus(function (status) {
                        if (status) {
                            var name = naverLogin.user.getName();

                            if( name === undefined || name == null) {
                                alert("이름은 필수정보입니다. 정보제공을 동의해주세요.");
                                naverLogin.reprompt();
                            }

                            var result = {}
                            result['name'] = name;
                            result['uniqId'] = naverLogin.user.getId();
                            store.dispatch('loginWithNaver', result)
                                .then(() => {
                                    //네이버는 리다이렉트 페이지 프레그먼트 형태로 바로 넘어가서 안 먹힘
                                })
                        } else {
                            alert("AccssToken이 올바르지 않아 callback 처리에 실패하였습니다.");
                            window.location.replace('/login')
                        }
                    });
                });
            }
        },
        mounted() {
            const script = document.createElement('script')
            script.src = 'https://static.nid.naver.com/js/naveridlogin_js_sdk_2.0.0.js'
            script.onload = () => this.naverLogin(this.$store)
            document.body.appendChild(script)
        },
    }

</script>

<style scoped>

</style>