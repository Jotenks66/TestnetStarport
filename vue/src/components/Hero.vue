<template>
  <div class="sp-component sp-welcome">
    <div class="sp-box sp-shadow sp-welcome__main">
      <div class="sp-welcome__header">
        <h2>Welcome to Gaupachain!</h2>
      </div>
      <p>
        This Blockchain is built and maintained by Gaupa Labs.<br />
        Find the main chain repository and a blockchain client below.
        <br /><br />
        Enhancing the Supply Chain of today with the Blockchain of tomorrow.
      </p>
      <div class="sp-welcome__btns">
        <button
          href="https://github.com/GaupaLabs/gaupachain"
          target="_blank"
          type="secondary"
          class="sp-welcome__btns__github"
          >Chain Repository</button>
          <button
          href="https://github.com/GaupaLabs/gaupachainclient"
          target="_blank"
          type="secondary"
          class="sp-welcome__btns__github"
          >Chain Client</button>
      </div>
    </div>
  </div>
</template>

<script>
export default {
	name: 'Hero',
	components: {
  },
}
</script>