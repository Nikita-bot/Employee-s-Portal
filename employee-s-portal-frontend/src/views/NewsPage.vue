<template>
    <div id="app">
        <HeaderComponent />
        <div class="container">
            <SideMenu />
            <NewsComponent/>
        </div>
    </div>    
</template>

<script setup>
    import HeaderComponent from '@/components/HadderComponent.vue';
    import NewsComponent from '@/components/NewsComponent.vue';
    import SideMenu from '@/components/SideMenu.vue';
</script>

