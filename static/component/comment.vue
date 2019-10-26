<template>
   <hello>hello</hello>
</template>

<hello></hello>
<script type="application/javascript">
   import hello from 'static/component/comment.vue'
    Vue.component('my-component-name', {
        // ... 选项 ...
        template: hello,
    })
</script>
