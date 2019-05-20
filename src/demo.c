        
int main()
{
    printf("string in main().\n");
    return 0;
}

void x_entry()
{
    printf("string in x_entry().\n");

    int ret = 11;
    ret = main();
    exit(ret);
}

