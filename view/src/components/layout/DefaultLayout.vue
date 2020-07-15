<template>
    <div>
        <my-header/>
        <slot/>
        <my-footer/>
    </div>
</template>

<script>
    import MyHeader from "./MyHeader";
    import MyFooter from "./MyFooter";
    export default {
        components: {MyFooter, MyHeader}
    }
</script>