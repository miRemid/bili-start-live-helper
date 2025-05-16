<script>
    import { GetQRCode } from "../wailsjs/go/main/App.js";
    let url = ""
    function getQRcode() {
        GetQRCode().then((result) => (url = result.data.url)).catch((error) => {
            console.error("Error:", error);
        });
    }


</script>

<h3>Dogs API</h3>
<div>
    <button class="btn" on:click={getQRcode}>
        SDKjfkdsf
    </button>
    {url}
</div>
<br />

<style>
    .btn:focus {
        border-width: 3px;
    }
</style>
