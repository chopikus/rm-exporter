<script lang="ts">
    import {
        Alert,
        Button,
        P,
        Input,
        Label,
        Spinner,
        Footer,
        A,
    } from "flowbite-svelte";
    import {
        ArrowRightOutline,
        InfoCircleSolid,
        TabletSolid,
        CloseOutline,
    } from "flowbite-svelte-icons";
    import {
        ReadDocs,
        IsIpValid,
        GetAppVersion,
    } from "../../wailsjs/go/main/App.js";
    import { push } from "svelte-spa-router";
    import { BrowserOpenURL } from "../../wailsjs/runtime/runtime.js";

    const source = "https://github.com/chopikus/rm-exporter";
    let version = $state("");

    GetAppVersion().then((v: string) => {
        version = v;
    });

    let rmIp: string = $state("10.11.99.1");
    let isRmIpValid: boolean = $state(true);
    let loading: boolean = $state(false);
    let show_error: boolean = $state(false);
    let error_message: string = $state("");

    $effect(() => {
        if (rmIp) {
            IsIpValid(rmIp).then((result: boolean) => (isRmIpValid = result));
        }
    });

    function onNext() {
        if (loading) {
            return;
        }
        show_error = false;
        loading = true;
        ReadDocs(rmIp)
            .then((_: any) => push("/files"))
            .catch((err: Error) => {
                console.log(
                    "Couldn't connect to reMarkable tablet! Make sure the IP address is correct.",
                );
                error_message = err.toString();
                show_error = true;
            })
            .finally(() => {
                loading = false;
            });
    }

    const onSourceClick = () => {
        BrowserOpenURL(source);
    };
</script>

{#key show_error}
    {#if show_error}
        <Alert
            color="red"
            dismissable={true}
            class="border-t-4 top-0 absolute left-1/2 transform -translate-x-1/2 w-full"
        >
            <div class="flex items-center gap-1">
                <InfoCircleSolid class="w-5 h-5" />
                <span class="text-lg font-medium"
                    >Couldn't connect to reMarkable!</span
                >
            </div>
            <p class="mt-2 mb-4 text-sm">{error_message}</p>
        </Alert>
    {/if}
{/key}

<main class="flex flex-col h-full justify-center items-center content-center">
    <P size="3xl">rm-exporter</P>
    <div
        class="flex flex-col flex-wrap content-center items-center justify-between h-44"
    >
        <Label class="space-y-3 m-5">
            <span>Enter the IP address of your reMarkable tablet:</span>
            <Input
                placeholder=""
                size="sm"
                bind:value={rmIp}
                color={isRmIpValid ? "base" : "red"}
            >
                <TabletSolid slot="left" class="w-4 h-6" />
            </Input>
        </Label>

        <Button pill size="xl" on:click={onNext}>
            {#if !loading}
                <ArrowRightOutline class="w-5 h-5 me-2" /> Next
            {:else}
                <Spinner />
            {/if}
        </Button>
    </div>
</main>
<Footer class="absolute bottom-0">
    <P size="sm"
        >Version {version}. Source available at <A onclick={onSourceClick}
            >{source}</A
        ></P
    >
</Footer>
