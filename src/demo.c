+-...+int main()
{
    printf("This is a xc demo.\n");
    return 0;
}

void xc_entry()
{
    int ret = 11;
    ret = main();
    exit(ret);
}

