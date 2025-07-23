<template>
    <div id="app">
        <HeaderComponent />
        <div class="container">
            <SideMenu />
            <SupportComponent />
        </div>
    </div>    
</template>

<script setup>
    import HeaderComponent from '@/components/HadderComponent.vue';
    import SupportComponent from '@/components/SupportComponent.vue';
    import SideMenu from '@/components/SideMenu.vue';
</script>