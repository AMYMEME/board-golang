<template>
    <a id="google-login-btn" type="button" style="display:inline-flex; align-items:center; min-height:50px; width: 222px;
      background-color:#4285F4; font-size:16px; color:white; text-decoration:none; border-radius:5px">
        <div style="background-color: white; margin:2px; padding-top:6px; padding-bottom:6px;border-radius:5px">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 533.5 544.3"
                 width="52px" height="35" style="display:inline-flex; align-items:center;" >
                <path d="M533.5 278.4c0-18.5-1.5-37.1-4.7-55.3H272.1v104.8h147c-6.1 33.8-25.7 63.7-54.4 82.7v68h87.7c51.5-47.4 81.1-117.4 81.1-200.2z" fill="#4285f4"/>
                <path d="M272.1 544.3c73.4 0 135.3-24.1 180.4-65.7l-87.7-68c-24.4 16.6-55.9 26-92.6 26-71 0-131.2-47.9-152.8-112.3H28.9v70.1c46.2 91.9 140.3 149.9 243.2 149.9z" fill="#34a853"/>
                <path d="M119.3 324.3c-11.4-33.8-11.4-70.4 0-104.2V150H28.9c-38.6 76.9-38.6 167.5 0 244.4l90.4-70.1z" fill="#fbbc04"/>
                <path d="M272.1 107.7c38.8-.6 76.3 14 104.4 40.8l77.7-77.7C405 24.6 339.7-.8 272.1 0 169.2 0 75.1 58 28.9 150l90.4 70.1c21.5-64.5 81.8-112.4 152.8-112.4z" fill="#ea4335"/>
            </svg>
        </div>
        <div style="padding-left: 12px">
            구글 계정으로 로그인
        </div>
    </a>
</template>

<script>
    export default {
        name: "AuthGoogle",
        methods: {
            googleLogin(store) {
                var GoogleAuth
                gapi.load('client:auth2', function () {
                    gapi.client.init({
                        'apiKey': 'AIzaSyB3ZjdYl2NRYTrSlZAY8T0E0GHZln_xFPg',
                        'clientId': '98250998013-tbbgns4o192oci849p3e8pf7ntfdgl8g.apps.googleusercontent.com',
                        'scope': 'https://www.googleapis.com/auth/drive.metadata.readonly',
                        'discoveryDocs': ["https://people.googleapis.com/$discovery/rest?version=v1"]
                    }).then(function () {
                        GoogleAuth = gapi.auth2.getAuthInstance();
                        var googleLoginBtn = document.getElementById('google-login-btn');
                        googleLoginBtn.addEventListener('click', function () {
                            if (GoogleAuth.isSignedIn.get()) {
                                gapi.client.people.people.get({
                                    'resourceName': 'people/me',
                                    'requestMask.includeField': 'person.names'
                                }).then((res) => {
                                    store.dispatch('loginWithGoogle', res.result.names[0])
                                        .then(() =>  store.commit('backRedirect'))
                                })
                            } else {
                                GoogleAuth.signIn();
                            }
                        })
                    })
                })
            },
        },
        mounted() {
            const script = document.createElement('script')
            script.src = 'https://apis.google.com/js/api.js'
            script.onload = (() => {
                this.googleLogin(this.$store)
            });
            document.body.appendChild(script)
        }
    }

</script>

<style scoped>

</style>