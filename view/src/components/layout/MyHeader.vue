<template>
    <header>
        <div>
            <b-navbar type="dark" variant="dark">
                <b-button v-b-toggle.sidebar-backdrop>
                    <span class="navbar-toggler-icon"></span>
                </b-button>

                <b-sidebar
                        id="sidebar-backdrop"
                        title="사이드바"
                        backdrop>
                    <div class="px-3 py-2">

                    </div>
                </b-sidebar>

                <b-navbar-nav class="ml-auto" right>
                    <b-nav-item href="/">Home</b-nav-item>
                    <b-nav-item href="/posts">Posts</b-nav-item>

                    <b-nav-item-dropdown text="User" right>
                        <b-dropdown-item href="#">Account</b-dropdown-item>
                        <b-dropdown-item href="#">Settings</b-dropdown-item>
                        <b-dropdown-item href="#">Logout></b-dropdown-item>
                    </b-nav-item-dropdown>
                </b-navbar-nav>
            </b-navbar>
        </div>
    </header>
</template>

<script>
    export default {
        name: "MyHeader"
    }
</script>

<style scoped>

</style>