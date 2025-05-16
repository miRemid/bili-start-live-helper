<script>
    import { GetQRCode } from "../wailsjs/go/main/App.js";
    import {EventsOn} from "../wailsjs/runtime/runtime.js";
    let url = ""
    function getQRcode() {
        GetQRCode().then((result) => (url = result.data.url)).catch((error) => {
            // Trigger an alert toast if the QR code generation fails
            console.error("Error:", error);
        });
    }
    EventsOn("qrCode", (event) => {
        // Login success, redirect to admin panel
    });
</script>

<div class="container">
    <button class="btn" on:click={getQRcode} >Login</button>
    <span style="color: black">{url}</span>
</div>

<style>
</style>
