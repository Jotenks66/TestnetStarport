<template>
	<div>
		<div class="container">
			<Hero />
			<SpTokenSend :address="address" refresh="true" />
			<SpTransferList :address="address" />
		</div>
	</div>
</template>

<script>
import Hero from '../components/Hero.vue'
export default {
	name: 'Index',
	computed: {
		address() {
			return this.$store.getters['common/wallet/address']
		}
	},
	components: { Hero },
}
</script>
