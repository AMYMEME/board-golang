<template>
    <div id="naverIdLogin" class="mt-4"></div>
</template>

<script>
    export default {
        name: "AuthNaver",
        data() {
            return {
                initiate: () => {
                    const naverLogin = new naver.LoginWithNaverId(
                        {
                            clientId: `_HXiSkrZSfB9vq94TsZC`,
                            callbackUrl: `http://localhost:8080/auth/naver`,
                            isPopup: false,
                            loginButton: {color: "green", type: 3, height: 50}
                        }
                    )
                    naverLogin.init()
                }
            }
        },
        mounted() {
            const script = document.createElement('script')
            script.src = 'https://static.nid.naver.com/js/naveridlogin_js_sdk_2.0.0.js'
            script.onload = () => this.initiate(this)
            document.body.appendChild(script)
        },
    }

</script>

<style scoped>

</style>