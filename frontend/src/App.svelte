<script lang="ts">
    import "./app.css";
    import Menu from "./pages/Menu.svelte";
    import FileSelection from "./pages/FileSelection.svelte";
    import Router from "svelte-spa-router";
    import ExportConfirmation from "./pages/ExportConfirmation.svelte";
    import Export from "./pages/Export.svelte";

    const routes = {
        "/": Menu,
        "/files": FileSelection,
        "/export-confirmation": ExportConfirmation,
        "/export": Export,
        "*": Menu,
    };
</script>

<Router {routes} />
